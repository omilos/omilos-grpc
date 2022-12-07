/**
 * @fileoverview gRPC-Web generated client stub for omilos_grpc
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_wrappers_pb = require('google-protobuf/google/protobuf/wrappers_pb.js')
const proto = {};
proto.omilos_grpc = require('./omilos_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.omilos_grpc.OmilosClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.omilos_grpc.OmilosPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.omilos_grpc.CastsRequest,
 *   !proto.omilos_grpc.CastsResponse>}
 */
const methodDescriptor_Omilos_Casts = new grpc.web.MethodDescriptor(
  '/omilos_grpc.Omilos/Casts',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.omilos_grpc.CastsRequest,
  proto.omilos_grpc.CastsResponse,
  /**
   * @param {!proto.omilos_grpc.CastsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.omilos_grpc.CastsResponse.deserializeBinary
);


/**
 * @param {!proto.omilos_grpc.CastsRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.omilos_grpc.CastsResponse>}
 *     The XHR Node Readable Stream
 */
proto.omilos_grpc.OmilosClient.prototype.casts =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/omilos_grpc.Omilos/Casts',
      request,
      metadata || {},
      methodDescriptor_Omilos_Casts);
};


/**
 * @param {!proto.omilos_grpc.CastsRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.omilos_grpc.CastsResponse>}
 *     The XHR Node Readable Stream
 */
proto.omilos_grpc.OmilosPromiseClient.prototype.casts =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/omilos_grpc.Omilos/Casts',
      request,
      metadata || {},
      methodDescriptor_Omilos_Casts);
};


module.exports = proto.omilos_grpc;


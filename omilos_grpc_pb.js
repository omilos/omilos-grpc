// GENERATED CODE -- DO NOT EDIT!

// Original file comments:
// User wants to fetch certificates, and credits
'use strict';
var grpc = require('@grpc/grpc-js');
var omilos_pb = require('./omilos_pb.js');
var google_protobuf_wrappers_pb = require('google-protobuf/google/protobuf/wrappers_pb.js');

function serialize_omilos_grpc_CastsRequest(arg) {
  if (!(arg instanceof omilos_pb.CastsRequest)) {
    throw new Error('Expected argument of type omilos_grpc.CastsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_omilos_grpc_CastsRequest(buffer_arg) {
  return omilos_pb.CastsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_omilos_grpc_CastsResponse(arg) {
  if (!(arg instanceof omilos_pb.CastsResponse)) {
    throw new Error('Expected argument of type omilos_grpc.CastsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_omilos_grpc_CastsResponse(buffer_arg) {
  return omilos_pb.CastsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var OmilosService = exports.OmilosService = {
  casts: {
    path: '/omilos_grpc.Omilos/Casts',
    requestStream: false,
    responseStream: true,
    requestType: omilos_pb.CastsRequest,
    responseType: omilos_pb.CastsResponse,
    requestSerialize: serialize_omilos_grpc_CastsRequest,
    requestDeserialize: deserialize_omilos_grpc_CastsRequest,
    responseSerialize: serialize_omilos_grpc_CastsResponse,
    responseDeserialize: deserialize_omilos_grpc_CastsResponse,
  },
};

exports.OmilosClient = grpc.makeGenericClientConstructor(OmilosService);

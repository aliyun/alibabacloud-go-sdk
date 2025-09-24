// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InitializeResponseBody
	GetCode() *string
	SetMessage(v string) *InitializeResponseBody
	GetMessage() *string
	SetRequestId(v string) *InitializeResponseBody
	GetRequestId() *string
	SetResult(v *InitializeResponseBodyResult) *InitializeResponseBody
	GetResult() *InitializeResponseBodyResult
}

type InitializeResponseBody struct {
	// Return code
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Return message
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4EB35****87EBA1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result
	Result *InitializeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s InitializeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitializeResponseBody) GoString() string {
	return s.String()
}

func (s *InitializeResponseBody) GetCode() *string {
	return s.Code
}

func (s *InitializeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InitializeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitializeResponseBody) GetResult() *InitializeResponseBodyResult {
	return s.Result
}

func (s *InitializeResponseBody) SetCode(v string) *InitializeResponseBody {
	s.Code = &v
	return s
}

func (s *InitializeResponseBody) SetMessage(v string) *InitializeResponseBody {
	s.Message = &v
	return s
}

func (s *InitializeResponseBody) SetRequestId(v string) *InitializeResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitializeResponseBody) SetResult(v *InitializeResponseBodyResult) *InitializeResponseBody {
	s.Result = v
	return s
}

func (s *InitializeResponseBody) Validate() error {
	return dara.Validate(s)
}

type InitializeResponseBodyResult struct {
	// Client configuration
	//
	// example:
	//
	// ***
	ClientCfg *string `json:"ClientCfg,omitempty" xml:"ClientCfg,omitempty"`
	// Standard encryption protocol for authentication.
	//
	// > This field is required when integrating with H5 web pages using iframe embedding.
	//
	// example:
	//
	// hksb7ba1b28130d24e015d*********
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// Authentication ID
	//
	// example:
	//
	// 08573be80f944d95ac812e019e3655a8
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
	// Web authentication URL
	//
	// example:
	//
	// http****
	TransactionUrl *string `json:"TransactionUrl,omitempty" xml:"TransactionUrl,omitempty"`
}

func (s InitializeResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s InitializeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *InitializeResponseBodyResult) GetClientCfg() *string {
	return s.ClientCfg
}

func (s *InitializeResponseBodyResult) GetProtocol() *string {
	return s.Protocol
}

func (s *InitializeResponseBodyResult) GetTransactionId() *string {
	return s.TransactionId
}

func (s *InitializeResponseBodyResult) GetTransactionUrl() *string {
	return s.TransactionUrl
}

func (s *InitializeResponseBodyResult) SetClientCfg(v string) *InitializeResponseBodyResult {
	s.ClientCfg = &v
	return s
}

func (s *InitializeResponseBodyResult) SetProtocol(v string) *InitializeResponseBodyResult {
	s.Protocol = &v
	return s
}

func (s *InitializeResponseBodyResult) SetTransactionId(v string) *InitializeResponseBodyResult {
	s.TransactionId = &v
	return s
}

func (s *InitializeResponseBodyResult) SetTransactionUrl(v string) *InitializeResponseBodyResult {
	s.TransactionUrl = &v
	return s
}

func (s *InitializeResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

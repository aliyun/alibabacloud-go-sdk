// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStopGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteStopGatewayResponseBody
	GetCode() *string
	SetData(v string) *DeleteStopGatewayResponseBody
	GetData() *string
	SetMessage(v string) *DeleteStopGatewayResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteStopGatewayResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeleteStopGatewayResponseBody
	GetSuccess() *string
}

type DeleteStopGatewayResponseBody struct {
	// The returned status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The status code indicating the result of the deletion:
	//
	// - **0**: Success. The metadata was deleted.
	//
	// - **-1**: A system error occurred.
	//
	// - **-2**: The specified database gateway does not exist.
	//
	// - **-3**: The database gateway is still active (not stopped) and its metadata cannot be deleted.
	//
	// - **-4**: Failed to delete the metadata.
	//
	// example:
	//
	// 0
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// > When the request is successful, this parameter returns **Successful**. When the request fails, this parameter returns exception information such as error codes.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FC6C0929-29E1-59FD-8DFE-70D9D41E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful:
	//
	// - **true**: The operation is successful.
	//
	// - **false**: The operation failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteStopGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStopGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStopGatewayResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteStopGatewayResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteStopGatewayResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteStopGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteStopGatewayResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteStopGatewayResponseBody) SetCode(v string) *DeleteStopGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteStopGatewayResponseBody) SetData(v string) *DeleteStopGatewayResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteStopGatewayResponseBody) SetMessage(v string) *DeleteStopGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteStopGatewayResponseBody) SetRequestId(v string) *DeleteStopGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStopGatewayResponseBody) SetSuccess(v string) *DeleteStopGatewayResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteStopGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}

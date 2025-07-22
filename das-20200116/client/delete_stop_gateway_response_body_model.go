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
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of the DeleteStopGateway operation. Valid values:
	//
	// 	- **0**: The metadata of the DBGateway is deleted.
	//
	// 	- **-1**: A system error occurs.
	//
	// 	- **-2**: The DBGateway does not exist.
	//
	// 	- **-3**: The DBGateway is not stopped and the metadata cannot be deleted.
	//
	// 	- **-4**: The metadata of the DBGateway fails to be deleted.
	//
	// example:
	//
	// 0
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
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
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
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

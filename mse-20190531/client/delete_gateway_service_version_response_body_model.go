// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayServiceVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteGatewayServiceVersionResponseBody
	GetCode() *int32
	SetData(v int64) *DeleteGatewayServiceVersionResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *DeleteGatewayServiceVersionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteGatewayServiceVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteGatewayServiceVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteGatewayServiceVersionResponseBody
	GetSuccess() *bool
}

type DeleteGatewayServiceVersionResponseBody struct {
	// The response code returned.
	//
	// example:
	//
	// 1
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// 1
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// 	- If the request is successful, a success message is returned.
	//
	// 	- If the request fails, an error message is returned, such as the "TaskId not found" message.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 29D52777-BE96-563E-BC6B-796ACE47A7A5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteGatewayServiceVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayServiceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayServiceVersionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteGatewayServiceVersionResponseBody) GetData() *int64 {
	return s.Data
}

func (s *DeleteGatewayServiceVersionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteGatewayServiceVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteGatewayServiceVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGatewayServiceVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteGatewayServiceVersionResponseBody) SetCode(v int32) *DeleteGatewayServiceVersionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGatewayServiceVersionResponseBody) SetData(v int64) *DeleteGatewayServiceVersionResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteGatewayServiceVersionResponseBody) SetHttpStatusCode(v int32) *DeleteGatewayServiceVersionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteGatewayServiceVersionResponseBody) SetMessage(v string) *DeleteGatewayServiceVersionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGatewayServiceVersionResponseBody) SetRequestId(v string) *DeleteGatewayServiceVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayServiceVersionResponseBody) SetSuccess(v bool) *DeleteGatewayServiceVersionResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGatewayServiceVersionResponseBody) Validate() error {
	return dara.Validate(s)
}

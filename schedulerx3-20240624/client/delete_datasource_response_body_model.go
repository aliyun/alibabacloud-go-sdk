// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteDatasourceResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteDatasourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDatasourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDatasourceResponseBody
	GetSuccess() *bool
}

type DeleteDatasourceResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CF99C381-C8F6-5A8D-8C24-57F46B706D2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDatasourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatasourceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteDatasourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDatasourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDatasourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDatasourceResponseBody) SetCode(v int32) *DeleteDatasourceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDatasourceResponseBody) SetMessage(v string) *DeleteDatasourceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDatasourceResponseBody) SetRequestId(v string) *DeleteDatasourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDatasourceResponseBody) SetSuccess(v bool) *DeleteDatasourceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDatasourceResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRumAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteRumAppResponseBody
	GetCode() *int32
	SetHttpStatusCode(v int32) *DeleteRumAppResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteRumAppResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteRumAppResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DeleteRumAppResponseBody
	GetResourceGroupId() *string
	SetResult(v string) *DeleteRumAppResponseBody
	GetResult() *string
	SetSuccess(v bool) *DeleteRumAppResponseBody
	GetSuccess() *bool
}

type DeleteRumAppResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4C518054-852F-4023-ABC1-4AF95FF7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2eq4peca****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The message that appears when the application is deleted.
	//
	// example:
	//
	// Success to delete app.
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRumAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRumAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRumAppResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteRumAppResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteRumAppResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteRumAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRumAppResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteRumAppResponseBody) GetResult() *string {
	return s.Result
}

func (s *DeleteRumAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteRumAppResponseBody) SetCode(v int32) *DeleteRumAppResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRumAppResponseBody) SetHttpStatusCode(v int32) *DeleteRumAppResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteRumAppResponseBody) SetMessage(v string) *DeleteRumAppResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteRumAppResponseBody) SetRequestId(v string) *DeleteRumAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRumAppResponseBody) SetResourceGroupId(v string) *DeleteRumAppResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteRumAppResponseBody) SetResult(v string) *DeleteRumAppResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteRumAppResponseBody) SetSuccess(v bool) *DeleteRumAppResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteRumAppResponseBody) Validate() error {
	return dara.Validate(s)
}

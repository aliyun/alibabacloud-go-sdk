// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResourceManagerResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ChangeResourceManagerResourceGroupResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *ChangeResourceManagerResourceGroupResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ChangeResourceManagerResourceGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChangeResourceManagerResourceGroupResponseBody
	GetSuccess() *bool
}

type ChangeResourceManagerResourceGroupResponseBody struct {
	// Indicates whether the resource group is changed.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AFAE64E-D1BE-432B-A9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeResourceManagerResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceManagerResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceManagerResourceGroupResponseBody) GetData() *bool {
	return s.Data
}

func (s *ChangeResourceManagerResourceGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ChangeResourceManagerResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeResourceManagerResourceGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeResourceManagerResourceGroupResponseBody) SetData(v bool) *ChangeResourceManagerResourceGroupResponseBody {
	s.Data = &v
	return s
}

func (s *ChangeResourceManagerResourceGroupResponseBody) SetHttpStatusCode(v int32) *ChangeResourceManagerResourceGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ChangeResourceManagerResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceManagerResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeResourceManagerResourceGroupResponseBody) SetSuccess(v bool) *ChangeResourceManagerResourceGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeResourceManagerResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

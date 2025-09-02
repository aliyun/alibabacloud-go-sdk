// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *CreateProjectResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *CreateProjectResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateProjectResponseBody
	GetSuccess() *bool
}

type CreateProjectResponseBody struct {
	// The workspace ID.
	//
	// example:
	//
	// 466230
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateProjectResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateProjectResponseBody) SetData(v int64) *CreateProjectResponseBody {
	s.Data = &v
	return s
}

func (s *CreateProjectResponseBody) SetHttpStatusCode(v int32) *CreateProjectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProjectResponseBody) SetSuccess(v bool) *CreateProjectResponseBody {
	s.Success = &v
	return s
}

func (s *CreateProjectResponseBody) Validate() error {
	return dara.Validate(s)
}

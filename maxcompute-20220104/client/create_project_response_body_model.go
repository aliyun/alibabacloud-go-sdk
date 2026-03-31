// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *CreateProjectResponseBody
	GetData() *string
	SetRequestId(v string) *CreateProjectResponseBody
	GetRequestId() *string
}

type CreateProjectResponseBody struct {
	// The returned result.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0b87b7a316654730544735643e9200
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProjectResponseBody) SetData(v string) *CreateProjectResponseBody {
	s.Data = &v
	return s
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProjectResponseBody) Validate() error {
	return dara.Validate(s)
}

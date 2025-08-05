// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectIpWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *UpdateProjectIpWhiteListResponseBody
	GetData() *string
	SetRequestId(v string) *UpdateProjectIpWhiteListResponseBody
	GetRequestId() *string
}

type UpdateProjectIpWhiteListResponseBody struct {
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
	// 0bc12e4316675560945192024e1044
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateProjectIpWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectIpWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectIpWhiteListResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateProjectIpWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProjectIpWhiteListResponseBody) SetData(v string) *UpdateProjectIpWhiteListResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateProjectIpWhiteListResponseBody) SetRequestId(v string) *UpdateProjectIpWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProjectIpWhiteListResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVodTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateVodTemplateResponseBody
	GetRequestId() *string
	SetVodTemplateId(v string) *UpdateVodTemplateResponseBody
	GetVodTemplateId() *string
}

type UpdateVodTemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A*****F6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the snapshot template.
	//
	// example:
	//
	// 8c75a02e339b*****0b0d2c48171a22
	VodTemplateId *string `json:"VodTemplateId,omitempty" xml:"VodTemplateId,omitempty"`
}

func (s UpdateVodTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVodTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVodTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVodTemplateResponseBody) GetVodTemplateId() *string {
	return s.VodTemplateId
}

func (s *UpdateVodTemplateResponseBody) SetRequestId(v string) *UpdateVodTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVodTemplateResponseBody) SetVodTemplateId(v string) *UpdateVodTemplateResponseBody {
	s.VodTemplateId = &v
	return s
}

func (s *UpdateVodTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

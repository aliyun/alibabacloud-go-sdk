// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVodTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVodTemplateResponseBody
	GetRequestId() *string
	SetVodTemplateId(v string) *DeleteVodTemplateResponseBody
	GetVodTemplateId() *string
}

type DeleteVodTemplateResponseBody struct {
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
	// f5b228fe6930e*****d6bf55bd87789
	VodTemplateId *string `json:"VodTemplateId,omitempty" xml:"VodTemplateId,omitempty"`
}

func (s DeleteVodTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVodTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVodTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVodTemplateResponseBody) GetVodTemplateId() *string {
	return s.VodTemplateId
}

func (s *DeleteVodTemplateResponseBody) SetRequestId(v string) *DeleteVodTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVodTemplateResponseBody) SetVodTemplateId(v string) *DeleteVodTemplateResponseBody {
	s.VodTemplateId = &v
	return s
}

func (s *DeleteVodTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

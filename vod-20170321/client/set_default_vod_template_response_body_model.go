// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultVodTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDefaultVodTemplateResponseBody
	GetRequestId() *string
	SetVodTemplateId(v string) *SetDefaultVodTemplateResponseBody
	GetVodTemplateId() *string
}

type SetDefaultVodTemplateResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VodTemplateId *string `json:"VodTemplateId,omitempty" xml:"VodTemplateId,omitempty"`
}

func (s SetDefaultVodTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultVodTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *SetDefaultVodTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDefaultVodTemplateResponseBody) GetVodTemplateId() *string {
	return s.VodTemplateId
}

func (s *SetDefaultVodTemplateResponseBody) SetRequestId(v string) *SetDefaultVodTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDefaultVodTemplateResponseBody) SetVodTemplateId(v string) *SetDefaultVodTemplateResponseBody {
	s.VodTemplateId = &v
	return s
}

func (s *SetDefaultVodTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

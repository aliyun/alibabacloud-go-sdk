// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreateDataQualityTemplateResponseBody
	GetId() *string
	SetRequestId(v string) *CreateDataQualityTemplateResponseBody
	GetRequestId() *string
}

type CreateDataQualityTemplateResponseBody struct {
	// example:
	//
	// USER-DEFINED:2001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 0bc14115***159376359
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataQualityTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataQualityTemplateResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateDataQualityTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataQualityTemplateResponseBody) SetId(v string) *CreateDataQualityTemplateResponseBody {
	s.Id = &v
	return s
}

func (s *CreateDataQualityTemplateResponseBody) SetRequestId(v string) *CreateDataQualityTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataQualityTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

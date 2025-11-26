// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ApiTemplate) *GetApiTemplateResponseBody
	GetData() *ApiTemplate
	SetRequestId(v string) *GetApiTemplateResponseBody
	GetRequestId() *string
}

type GetApiTemplateResponseBody struct {
	// Deprecated
	//
	// The content of the API operation template.
	Data *ApiTemplate `json:"Data,omitempty" xml:"Data,omitempty"`
	// Request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApiTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApiTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetApiTemplateResponseBody) GetData() *ApiTemplate {
	return s.Data
}

func (s *GetApiTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApiTemplateResponseBody) SetData(v *ApiTemplate) *GetApiTemplateResponseBody {
	s.Data = v
	return s
}

func (s *GetApiTemplateResponseBody) SetRequestId(v string) *GetApiTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApiTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

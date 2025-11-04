// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContentAnalyzeConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContentAnalyzeConfig(v *GetContentAnalyzeConfigResponseBodyContentAnalyzeConfig) *GetContentAnalyzeConfigResponseBody
	GetContentAnalyzeConfig() *GetContentAnalyzeConfigResponseBodyContentAnalyzeConfig
	SetRequestId(v string) *GetContentAnalyzeConfigResponseBody
	GetRequestId() *string
}

type GetContentAnalyzeConfigResponseBody struct {
	ContentAnalyzeConfig *GetContentAnalyzeConfigResponseBodyContentAnalyzeConfig `json:"ContentAnalyzeConfig,omitempty" xml:"ContentAnalyzeConfig,omitempty" type:"Struct"`
	// example:
	//
	// 31FEC819-2344-5771-9366-9172DB0D26C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetContentAnalyzeConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetContentAnalyzeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetContentAnalyzeConfigResponseBody) GetContentAnalyzeConfig() *GetContentAnalyzeConfigResponseBodyContentAnalyzeConfig {
	return s.ContentAnalyzeConfig
}

func (s *GetContentAnalyzeConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetContentAnalyzeConfigResponseBody) SetContentAnalyzeConfig(v *GetContentAnalyzeConfigResponseBodyContentAnalyzeConfig) *GetContentAnalyzeConfigResponseBody {
	s.ContentAnalyzeConfig = v
	return s
}

func (s *GetContentAnalyzeConfigResponseBody) SetRequestId(v string) *GetContentAnalyzeConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContentAnalyzeConfigResponseBody) Validate() error {
	if s.ContentAnalyzeConfig != nil {
		if err := s.ContentAnalyzeConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetContentAnalyzeConfigResponseBodyContentAnalyzeConfig struct {
	// example:
	//
	// true
	Auto *bool `json:"Auto,omitempty" xml:"Auto,omitempty"`
	// example:
	//
	// TEXT,FACE
	SaveType *string `json:"SaveType,omitempty" xml:"SaveType,omitempty"`
	// example:
	//
	// S00000101-100070
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetContentAnalyzeConfigResponseBodyContentAnalyzeConfig) String() string {
	return dara.Prettify(s)
}

func (s GetContentAnalyzeConfigResponseBodyContentAnalyzeConfig) GoString() string {
	return s.String()
}

func (s *GetContentAnalyzeConfigResponseBodyContentAnalyzeConfig) GetAuto() *bool {
	return s.Auto
}

func (s *GetContentAnalyzeConfigResponseBodyContentAnalyzeConfig) GetSaveType() *string {
	return s.SaveType
}

func (s *GetContentAnalyzeConfigResponseBodyContentAnalyzeConfig) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetContentAnalyzeConfigResponseBodyContentAnalyzeConfig) SetAuto(v bool) *GetContentAnalyzeConfigResponseBodyContentAnalyzeConfig {
	s.Auto = &v
	return s
}

func (s *GetContentAnalyzeConfigResponseBodyContentAnalyzeConfig) SetSaveType(v string) *GetContentAnalyzeConfigResponseBodyContentAnalyzeConfig {
	s.SaveType = &v
	return s
}

func (s *GetContentAnalyzeConfigResponseBodyContentAnalyzeConfig) SetTemplateId(v string) *GetContentAnalyzeConfigResponseBodyContentAnalyzeConfig {
	s.TemplateId = &v
	return s
}

func (s *GetContentAnalyzeConfigResponseBodyContentAnalyzeConfig) Validate() error {
	return dara.Validate(s)
}

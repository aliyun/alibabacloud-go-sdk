// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageSceneLabelListConfResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetImageSceneLabelListConfResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetImageSceneLabelListConfResponse
	GetStatusCode() *int32
	SetBody(v *GetImageSceneLabelListConfResponseBody) *GetImageSceneLabelListConfResponse
	GetBody() *GetImageSceneLabelListConfResponseBody
}

type GetImageSceneLabelListConfResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageSceneLabelListConfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImageSceneLabelListConfResponse) String() string {
	return dara.Prettify(s)
}

func (s GetImageSceneLabelListConfResponse) GoString() string {
	return s.String()
}

func (s *GetImageSceneLabelListConfResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetImageSceneLabelListConfResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetImageSceneLabelListConfResponse) GetBody() *GetImageSceneLabelListConfResponseBody {
	return s.Body
}

func (s *GetImageSceneLabelListConfResponse) SetHeaders(v map[string]*string) *GetImageSceneLabelListConfResponse {
	s.Headers = v
	return s
}

func (s *GetImageSceneLabelListConfResponse) SetStatusCode(v int32) *GetImageSceneLabelListConfResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageSceneLabelListConfResponse) SetBody(v *GetImageSceneLabelListConfResponseBody) *GetImageSceneLabelListConfResponse {
	s.Body = v
	return s
}

func (s *GetImageSceneLabelListConfResponse) Validate() error {
	return dara.Validate(s)
}

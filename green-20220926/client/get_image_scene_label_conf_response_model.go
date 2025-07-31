// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageSceneLabelConfResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetImageSceneLabelConfResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetImageSceneLabelConfResponse
	GetStatusCode() *int32
	SetBody(v *GetImageSceneLabelConfResponseBody) *GetImageSceneLabelConfResponse
	GetBody() *GetImageSceneLabelConfResponseBody
}

type GetImageSceneLabelConfResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageSceneLabelConfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImageSceneLabelConfResponse) String() string {
	return dara.Prettify(s)
}

func (s GetImageSceneLabelConfResponse) GoString() string {
	return s.String()
}

func (s *GetImageSceneLabelConfResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetImageSceneLabelConfResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetImageSceneLabelConfResponse) GetBody() *GetImageSceneLabelConfResponseBody {
	return s.Body
}

func (s *GetImageSceneLabelConfResponse) SetHeaders(v map[string]*string) *GetImageSceneLabelConfResponse {
	s.Headers = v
	return s
}

func (s *GetImageSceneLabelConfResponse) SetStatusCode(v int32) *GetImageSceneLabelConfResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageSceneLabelConfResponse) SetBody(v *GetImageSceneLabelConfResponseBody) *GetImageSceneLabelConfResponse {
	s.Body = v
	return s
}

func (s *GetImageSceneLabelConfResponse) Validate() error {
	return dara.Validate(s)
}

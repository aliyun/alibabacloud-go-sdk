// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryIndustryLabelBagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryIndustryLabelBagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryIndustryLabelBagResponse
	GetStatusCode() *int32
	SetBody(v *QueryIndustryLabelBagResponseBody) *QueryIndustryLabelBagResponse
	GetBody() *QueryIndustryLabelBagResponseBody
}

type QueryIndustryLabelBagResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryIndustryLabelBagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryIndustryLabelBagResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryIndustryLabelBagResponse) GoString() string {
	return s.String()
}

func (s *QueryIndustryLabelBagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryIndustryLabelBagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryIndustryLabelBagResponse) GetBody() *QueryIndustryLabelBagResponseBody {
	return s.Body
}

func (s *QueryIndustryLabelBagResponse) SetHeaders(v map[string]*string) *QueryIndustryLabelBagResponse {
	s.Headers = v
	return s
}

func (s *QueryIndustryLabelBagResponse) SetStatusCode(v int32) *QueryIndustryLabelBagResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryIndustryLabelBagResponse) SetBody(v *QueryIndustryLabelBagResponseBody) *QueryIndustryLabelBagResponse {
	s.Body = v
	return s
}

func (s *QueryIndustryLabelBagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

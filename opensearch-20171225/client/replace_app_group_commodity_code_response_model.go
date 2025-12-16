// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceAppGroupCommodityCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReplaceAppGroupCommodityCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReplaceAppGroupCommodityCodeResponse
	GetStatusCode() *int32
	SetBody(v *ReplaceAppGroupCommodityCodeResponseBody) *ReplaceAppGroupCommodityCodeResponse
	GetBody() *ReplaceAppGroupCommodityCodeResponseBody
}

type ReplaceAppGroupCommodityCodeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReplaceAppGroupCommodityCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReplaceAppGroupCommodityCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s ReplaceAppGroupCommodityCodeResponse) GoString() string {
	return s.String()
}

func (s *ReplaceAppGroupCommodityCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReplaceAppGroupCommodityCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReplaceAppGroupCommodityCodeResponse) GetBody() *ReplaceAppGroupCommodityCodeResponseBody {
	return s.Body
}

func (s *ReplaceAppGroupCommodityCodeResponse) SetHeaders(v map[string]*string) *ReplaceAppGroupCommodityCodeResponse {
	s.Headers = v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponse) SetStatusCode(v int32) *ReplaceAppGroupCommodityCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponse) SetBody(v *ReplaceAppGroupCommodityCodeResponseBody) *ReplaceAppGroupCommodityCodeResponse {
	s.Body = v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

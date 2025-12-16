// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecondRankResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySecondRankResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySecondRankResponse
	GetStatusCode() *int32
	SetBody(v *ModifySecondRankResponseBody) *ModifySecondRankResponse
	GetBody() *ModifySecondRankResponseBody
}

type ModifySecondRankResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySecondRankResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySecondRankResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySecondRankResponse) GoString() string {
	return s.String()
}

func (s *ModifySecondRankResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySecondRankResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySecondRankResponse) GetBody() *ModifySecondRankResponseBody {
	return s.Body
}

func (s *ModifySecondRankResponse) SetHeaders(v map[string]*string) *ModifySecondRankResponse {
	s.Headers = v
	return s
}

func (s *ModifySecondRankResponse) SetStatusCode(v int32) *ModifySecondRankResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySecondRankResponse) SetBody(v *ModifySecondRankResponseBody) *ModifySecondRankResponse {
	s.Body = v
	return s
}

func (s *ModifySecondRankResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

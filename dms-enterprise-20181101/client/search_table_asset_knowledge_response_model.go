// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTableAssetKnowledgeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchTableAssetKnowledgeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchTableAssetKnowledgeResponse
	GetStatusCode() *int32
	SetBody(v *SearchTableAssetKnowledgeResponseBody) *SearchTableAssetKnowledgeResponse
	GetBody() *SearchTableAssetKnowledgeResponseBody
}

type SearchTableAssetKnowledgeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchTableAssetKnowledgeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchTableAssetKnowledgeResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchTableAssetKnowledgeResponse) GoString() string {
	return s.String()
}

func (s *SearchTableAssetKnowledgeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchTableAssetKnowledgeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchTableAssetKnowledgeResponse) GetBody() *SearchTableAssetKnowledgeResponseBody {
	return s.Body
}

func (s *SearchTableAssetKnowledgeResponse) SetHeaders(v map[string]*string) *SearchTableAssetKnowledgeResponse {
	s.Headers = v
	return s
}

func (s *SearchTableAssetKnowledgeResponse) SetStatusCode(v int32) *SearchTableAssetKnowledgeResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTableAssetKnowledgeResponse) SetBody(v *SearchTableAssetKnowledgeResponseBody) *SearchTableAssetKnowledgeResponse {
	s.Body = v
	return s
}

func (s *SearchTableAssetKnowledgeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

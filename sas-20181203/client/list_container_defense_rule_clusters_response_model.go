// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContainerDefenseRuleClustersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListContainerDefenseRuleClustersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListContainerDefenseRuleClustersResponse
	GetStatusCode() *int32
	SetBody(v *ListContainerDefenseRuleClustersResponseBody) *ListContainerDefenseRuleClustersResponse
	GetBody() *ListContainerDefenseRuleClustersResponseBody
}

type ListContainerDefenseRuleClustersResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListContainerDefenseRuleClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListContainerDefenseRuleClustersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListContainerDefenseRuleClustersResponse) GoString() string {
	return s.String()
}

func (s *ListContainerDefenseRuleClustersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListContainerDefenseRuleClustersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListContainerDefenseRuleClustersResponse) GetBody() *ListContainerDefenseRuleClustersResponseBody {
	return s.Body
}

func (s *ListContainerDefenseRuleClustersResponse) SetHeaders(v map[string]*string) *ListContainerDefenseRuleClustersResponse {
	s.Headers = v
	return s
}

func (s *ListContainerDefenseRuleClustersResponse) SetStatusCode(v int32) *ListContainerDefenseRuleClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContainerDefenseRuleClustersResponse) SetBody(v *ListContainerDefenseRuleClustersResponseBody) *ListContainerDefenseRuleClustersResponse {
	s.Body = v
	return s
}

func (s *ListContainerDefenseRuleClustersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

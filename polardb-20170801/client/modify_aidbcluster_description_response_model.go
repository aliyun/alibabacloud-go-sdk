// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAIDBClusterDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAIDBClusterDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAIDBClusterDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAIDBClusterDescriptionResponseBody) *ModifyAIDBClusterDescriptionResponse
	GetBody() *ModifyAIDBClusterDescriptionResponseBody
}

type ModifyAIDBClusterDescriptionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAIDBClusterDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAIDBClusterDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAIDBClusterDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyAIDBClusterDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAIDBClusterDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAIDBClusterDescriptionResponse) GetBody() *ModifyAIDBClusterDescriptionResponseBody {
	return s.Body
}

func (s *ModifyAIDBClusterDescriptionResponse) SetHeaders(v map[string]*string) *ModifyAIDBClusterDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyAIDBClusterDescriptionResponse) SetStatusCode(v int32) *ModifyAIDBClusterDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAIDBClusterDescriptionResponse) SetBody(v *ModifyAIDBClusterDescriptionResponseBody) *ModifyAIDBClusterDescriptionResponse {
	s.Body = v
	return s
}

func (s *ModifyAIDBClusterDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

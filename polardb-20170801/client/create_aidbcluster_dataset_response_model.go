// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAIDBClusterDatasetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAIDBClusterDatasetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAIDBClusterDatasetResponse
	GetStatusCode() *int32
	SetBody(v *CreateAIDBClusterDatasetResponseBody) *CreateAIDBClusterDatasetResponse
	GetBody() *CreateAIDBClusterDatasetResponseBody
}

type CreateAIDBClusterDatasetResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAIDBClusterDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAIDBClusterDatasetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAIDBClusterDatasetResponse) GoString() string {
	return s.String()
}

func (s *CreateAIDBClusterDatasetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAIDBClusterDatasetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAIDBClusterDatasetResponse) GetBody() *CreateAIDBClusterDatasetResponseBody {
	return s.Body
}

func (s *CreateAIDBClusterDatasetResponse) SetHeaders(v map[string]*string) *CreateAIDBClusterDatasetResponse {
	s.Headers = v
	return s
}

func (s *CreateAIDBClusterDatasetResponse) SetStatusCode(v int32) *CreateAIDBClusterDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAIDBClusterDatasetResponse) SetBody(v *CreateAIDBClusterDatasetResponseBody) *CreateAIDBClusterDatasetResponse {
	s.Body = v
	return s
}

func (s *CreateAIDBClusterDatasetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

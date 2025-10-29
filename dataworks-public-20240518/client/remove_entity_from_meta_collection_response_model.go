// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveEntityFromMetaCollectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveEntityFromMetaCollectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveEntityFromMetaCollectionResponse
	GetStatusCode() *int32
	SetBody(v *RemoveEntityFromMetaCollectionResponseBody) *RemoveEntityFromMetaCollectionResponse
	GetBody() *RemoveEntityFromMetaCollectionResponseBody
}

type RemoveEntityFromMetaCollectionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveEntityFromMetaCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveEntityFromMetaCollectionResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveEntityFromMetaCollectionResponse) GoString() string {
	return s.String()
}

func (s *RemoveEntityFromMetaCollectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveEntityFromMetaCollectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveEntityFromMetaCollectionResponse) GetBody() *RemoveEntityFromMetaCollectionResponseBody {
	return s.Body
}

func (s *RemoveEntityFromMetaCollectionResponse) SetHeaders(v map[string]*string) *RemoveEntityFromMetaCollectionResponse {
	s.Headers = v
	return s
}

func (s *RemoveEntityFromMetaCollectionResponse) SetStatusCode(v int32) *RemoveEntityFromMetaCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveEntityFromMetaCollectionResponse) SetBody(v *RemoveEntityFromMetaCollectionResponseBody) *RemoveEntityFromMetaCollectionResponse {
	s.Body = v
	return s
}

func (s *RemoveEntityFromMetaCollectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

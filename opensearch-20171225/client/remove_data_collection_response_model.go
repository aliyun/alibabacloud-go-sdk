// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDataCollectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveDataCollectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveDataCollectionResponse
	GetStatusCode() *int32
	SetBody(v *RemoveDataCollectionResponseBody) *RemoveDataCollectionResponse
	GetBody() *RemoveDataCollectionResponseBody
}

type RemoveDataCollectionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveDataCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveDataCollectionResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveDataCollectionResponse) GoString() string {
	return s.String()
}

func (s *RemoveDataCollectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveDataCollectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveDataCollectionResponse) GetBody() *RemoveDataCollectionResponseBody {
	return s.Body
}

func (s *RemoveDataCollectionResponse) SetHeaders(v map[string]*string) *RemoveDataCollectionResponse {
	s.Headers = v
	return s
}

func (s *RemoveDataCollectionResponse) SetStatusCode(v int32) *RemoveDataCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveDataCollectionResponse) SetBody(v *RemoveDataCollectionResponseBody) *RemoveDataCollectionResponse {
	s.Body = v
	return s
}

func (s *RemoveDataCollectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

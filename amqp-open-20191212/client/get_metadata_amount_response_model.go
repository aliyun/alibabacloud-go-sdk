// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetadataAmountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMetadataAmountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMetadataAmountResponse
	GetStatusCode() *int32
	SetBody(v *GetMetadataAmountResponseBody) *GetMetadataAmountResponse
	GetBody() *GetMetadataAmountResponseBody
}

type GetMetadataAmountResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetadataAmountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetadataAmountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMetadataAmountResponse) GoString() string {
	return s.String()
}

func (s *GetMetadataAmountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMetadataAmountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMetadataAmountResponse) GetBody() *GetMetadataAmountResponseBody {
	return s.Body
}

func (s *GetMetadataAmountResponse) SetHeaders(v map[string]*string) *GetMetadataAmountResponse {
	s.Headers = v
	return s
}

func (s *GetMetadataAmountResponse) SetStatusCode(v int32) *GetMetadataAmountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetadataAmountResponse) SetBody(v *GetMetadataAmountResponseBody) *GetMetadataAmountResponse {
	s.Body = v
	return s
}

func (s *GetMetadataAmountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

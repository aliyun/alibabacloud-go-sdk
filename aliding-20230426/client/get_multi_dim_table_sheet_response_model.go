// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiDimTableSheetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMultiDimTableSheetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMultiDimTableSheetResponse
	GetStatusCode() *int32
	SetBody(v *GetMultiDimTableSheetResponseBody) *GetMultiDimTableSheetResponse
	GetBody() *GetMultiDimTableSheetResponseBody
}

type GetMultiDimTableSheetResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMultiDimTableSheetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMultiDimTableSheetResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableSheetResponse) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableSheetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMultiDimTableSheetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMultiDimTableSheetResponse) GetBody() *GetMultiDimTableSheetResponseBody {
	return s.Body
}

func (s *GetMultiDimTableSheetResponse) SetHeaders(v map[string]*string) *GetMultiDimTableSheetResponse {
	s.Headers = v
	return s
}

func (s *GetMultiDimTableSheetResponse) SetStatusCode(v int32) *GetMultiDimTableSheetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultiDimTableSheetResponse) SetBody(v *GetMultiDimTableSheetResponseBody) *GetMultiDimTableSheetResponse {
	s.Body = v
	return s
}

func (s *GetMultiDimTableSheetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

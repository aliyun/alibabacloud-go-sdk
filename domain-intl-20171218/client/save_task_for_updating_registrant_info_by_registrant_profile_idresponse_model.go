// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse
	GetStatusCode() *int32
	SetBody(v *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponseBody) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse
	GetBody() *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponseBody
}

type SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse struct {
	Headers    map[string]*string                                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse) GoString() string {
	return s.String()
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse) GetBody() *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponseBody {
	return s.Body
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse) SetHeaders(v map[string]*string) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse {
	s.Headers = v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse) SetStatusCode(v int32) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse) SetBody(v *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponseBody) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse {
	s.Body = v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

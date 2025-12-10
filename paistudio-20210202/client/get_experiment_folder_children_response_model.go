// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentFolderChildrenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExperimentFolderChildrenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExperimentFolderChildrenResponse
	GetStatusCode() *int32
	SetBody(v *GetExperimentFolderChildrenResponseBody) *GetExperimentFolderChildrenResponse
	GetBody() *GetExperimentFolderChildrenResponseBody
}

type GetExperimentFolderChildrenResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExperimentFolderChildrenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExperimentFolderChildrenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentFolderChildrenResponse) GoString() string {
	return s.String()
}

func (s *GetExperimentFolderChildrenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExperimentFolderChildrenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExperimentFolderChildrenResponse) GetBody() *GetExperimentFolderChildrenResponseBody {
	return s.Body
}

func (s *GetExperimentFolderChildrenResponse) SetHeaders(v map[string]*string) *GetExperimentFolderChildrenResponse {
	s.Headers = v
	return s
}

func (s *GetExperimentFolderChildrenResponse) SetStatusCode(v int32) *GetExperimentFolderChildrenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExperimentFolderChildrenResponse) SetBody(v *GetExperimentFolderChildrenResponseBody) *GetExperimentFolderChildrenResponse {
	s.Body = v
	return s
}

func (s *GetExperimentFolderChildrenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

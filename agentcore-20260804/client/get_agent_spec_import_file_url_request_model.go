// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSpecImportFileUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentType(v string) *GetAgentSpecImportFileUrlRequest
	GetContentType() *string
}

type GetAgentSpecImportFileUrlRequest struct {
	// The Content-Type of the file to upload. The client must use the same value from the response when performing the PUT request.
	//
	// example:
	//
	// application/zip
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
}

func (s GetAgentSpecImportFileUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSpecImportFileUrlRequest) GoString() string {
	return s.String()
}

func (s *GetAgentSpecImportFileUrlRequest) GetContentType() *string {
	return s.ContentType
}

func (s *GetAgentSpecImportFileUrlRequest) SetContentType(v string) *GetAgentSpecImportFileUrlRequest {
	s.ContentType = &v
	return s
}

func (s *GetAgentSpecImportFileUrlRequest) Validate() error {
	return dara.Validate(s)
}

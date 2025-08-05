// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTempFileDownloadLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTempFileKey(v string) *GetTempFileDownloadLinkRequest
	GetTempFileKey() *string
}

type GetTempFileDownloadLinkRequest struct {
	// The key that is used to download a file.
	//
	// This parameter is required.
	//
	// example:
	//
	// temp/1797733170015112/report/r-000jdzknbp39cnf9hs99/r-000jdzknbp39cnf9hs99-total.csv
	TempFileKey *string `json:"TempFileKey,omitempty" xml:"TempFileKey,omitempty"`
}

func (s GetTempFileDownloadLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTempFileDownloadLinkRequest) GoString() string {
	return s.String()
}

func (s *GetTempFileDownloadLinkRequest) GetTempFileKey() *string {
	return s.TempFileKey
}

func (s *GetTempFileDownloadLinkRequest) SetTempFileKey(v string) *GetTempFileDownloadLinkRequest {
	s.TempFileKey = &v
	return s
}

func (s *GetTempFileDownloadLinkRequest) Validate() error {
	return dara.Validate(s)
}

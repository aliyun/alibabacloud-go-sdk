// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPreSignedUrlForPutObjectHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetPreSignedUrlForPutObjectHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetPreSignedUrlForPutObjectHeaders
	GetWorkspace() *string
}

type GetPreSignedUrlForPutObjectHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a9c3a20210af000
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetPreSignedUrlForPutObjectHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetPreSignedUrlForPutObjectHeaders) GoString() string {
	return s.String()
}

func (s *GetPreSignedUrlForPutObjectHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetPreSignedUrlForPutObjectHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetPreSignedUrlForPutObjectHeaders) SetCommonHeaders(v map[string]*string) *GetPreSignedUrlForPutObjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPreSignedUrlForPutObjectHeaders) SetWorkspace(v string) *GetPreSignedUrlForPutObjectHeaders {
	s.Workspace = &v
	return s
}

func (s *GetPreSignedUrlForPutObjectHeaders) Validate() error {
	return dara.Validate(s)
}

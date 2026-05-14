// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadFormInfo interface {
	dara.Model
	String() string
	GoString() string
	SetEndpoint(v string) *UploadFormInfo
	GetEndpoint() *string
	SetFormData(v map[string]*string) *UploadFormInfo
	GetFormData() map[string]*string
}

type UploadFormInfo struct {
	Endpoint *string            `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	FormData map[string]*string `json:"form_data,omitempty" xml:"form_data,omitempty"`
}

func (s UploadFormInfo) String() string {
	return dara.Prettify(s)
}

func (s UploadFormInfo) GoString() string {
	return s.String()
}

func (s *UploadFormInfo) GetEndpoint() *string {
	return s.Endpoint
}

func (s *UploadFormInfo) GetFormData() map[string]*string {
	return s.FormData
}

func (s *UploadFormInfo) SetEndpoint(v string) *UploadFormInfo {
	s.Endpoint = &v
	return s
}

func (s *UploadFormInfo) SetFormData(v map[string]*string) *UploadFormInfo {
	s.FormData = v
	return s
}

func (s *UploadFormInfo) Validate() error {
	return dara.Validate(s)
}

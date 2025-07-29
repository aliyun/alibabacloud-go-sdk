// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterDiagnosisResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *GetClusterDiagnosisResultRequest
	GetLanguage() *string
}

type GetClusterDiagnosisResultRequest struct {
	// The query language.
	//
	// example:
	//
	// zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s GetClusterDiagnosisResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClusterDiagnosisResultRequest) GoString() string {
	return s.String()
}

func (s *GetClusterDiagnosisResultRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetClusterDiagnosisResultRequest) SetLanguage(v string) *GetClusterDiagnosisResultRequest {
	s.Language = &v
	return s
}

func (s *GetClusterDiagnosisResultRequest) Validate() error {
	return dara.Validate(s)
}

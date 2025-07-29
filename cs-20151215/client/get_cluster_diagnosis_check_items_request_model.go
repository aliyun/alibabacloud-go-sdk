// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterDiagnosisCheckItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *GetClusterDiagnosisCheckItemsRequest
	GetLanguage() *string
}

type GetClusterDiagnosisCheckItemsRequest struct {
	// The query language.
	//
	// example:
	//
	// zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s GetClusterDiagnosisCheckItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClusterDiagnosisCheckItemsRequest) GoString() string {
	return s.String()
}

func (s *GetClusterDiagnosisCheckItemsRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetClusterDiagnosisCheckItemsRequest) SetLanguage(v string) *GetClusterDiagnosisCheckItemsRequest {
	s.Language = &v
	return s
}

func (s *GetClusterDiagnosisCheckItemsRequest) Validate() error {
	return dara.Validate(s)
}

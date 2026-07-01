// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlterSearchLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSearchLibConfig(v string) *AlterSearchLibRequest
	GetSearchLibConfig() *string
	SetSearchLibName(v string) *AlterSearchLibRequest
	GetSearchLibName() *string
}

type AlterSearchLibRequest struct {
	// The configuration of the search library, in JSON string format. The fields include the following:
	//
	// - faceGroupIds: The IDs of custom face recognition libraries. You can create these libraries by calling the CreateRecognitionLib operation. You can specify up to three library IDs, separated by commas (,).
	//
	// example:
	//
	// {"faceGroupIds":"xxx1,xxx2,xx3"}
	SearchLibConfig *string `json:"SearchLibConfig,omitempty" xml:"SearchLibConfig,omitempty"`
	// The search library.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	SearchLibName *string `json:"SearchLibName,omitempty" xml:"SearchLibName,omitempty"`
}

func (s AlterSearchLibRequest) String() string {
	return dara.Prettify(s)
}

func (s AlterSearchLibRequest) GoString() string {
	return s.String()
}

func (s *AlterSearchLibRequest) GetSearchLibConfig() *string {
	return s.SearchLibConfig
}

func (s *AlterSearchLibRequest) GetSearchLibName() *string {
	return s.SearchLibName
}

func (s *AlterSearchLibRequest) SetSearchLibConfig(v string) *AlterSearchLibRequest {
	s.SearchLibConfig = &v
	return s
}

func (s *AlterSearchLibRequest) SetSearchLibName(v string) *AlterSearchLibRequest {
	s.SearchLibName = &v
	return s
}

func (s *AlterSearchLibRequest) Validate() error {
	return dara.Validate(s)
}

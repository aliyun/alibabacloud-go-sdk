// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSearchLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSearchLibConfig(v string) *CreateSearchLibRequest
	GetSearchLibConfig() *string
	SetSearchLibName(v string) *CreateSearchLibRequest
	GetSearchLibName() *string
}

type CreateSearchLibRequest struct {
	// The configuration of the search library, in JSON string format. The fields include the following:
	//
	// - faceGroupIds: The IDs of custom face recognition libraries. You can create these libraries by calling the CreateRecognitionLib operation. You can specify up to three library IDs, separated by commas (,).
	//
	// example:
	//
	// {"faceGroupIds":"xxx1,xxx2,xx3"}
	SearchLibConfig *string `json:"SearchLibConfig,omitempty" xml:"SearchLibConfig,omitempty"`
	// The name of the search library. The name can contain letters and digits. For IP camera (IPC) monitoring scenarios, the name must start with the `IPCamera_` prefix. For other scenarios, you can use a custom name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	SearchLibName *string `json:"SearchLibName,omitempty" xml:"SearchLibName,omitempty"`
}

func (s CreateSearchLibRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSearchLibRequest) GoString() string {
	return s.String()
}

func (s *CreateSearchLibRequest) GetSearchLibConfig() *string {
	return s.SearchLibConfig
}

func (s *CreateSearchLibRequest) GetSearchLibName() *string {
	return s.SearchLibName
}

func (s *CreateSearchLibRequest) SetSearchLibConfig(v string) *CreateSearchLibRequest {
	s.SearchLibConfig = &v
	return s
}

func (s *CreateSearchLibRequest) SetSearchLibName(v string) *CreateSearchLibRequest {
	s.SearchLibName = &v
	return s
}

func (s *CreateSearchLibRequest) Validate() error {
	return dara.Validate(s)
}

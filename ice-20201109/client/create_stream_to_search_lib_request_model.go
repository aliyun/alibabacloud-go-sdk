// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStreamToSearchLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInput(v string) *CreateStreamToSearchLibRequest
	GetInput() *string
	SetNamespace(v string) *CreateStreamToSearchLibRequest
	GetNamespace() *string
	SetSearchLibName(v string) *CreateStreamToSearchLibRequest
	GetSearchLibName() *string
}

type CreateStreamToSearchLibRequest struct {
	// The URL of the live stream to be ingested and analyzed.
	//
	// example:
	//
	// rtmp://xxx
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The namespace.
	//
	// example:
	//
	// name-1
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The search library.
	//
	// example:
	//
	// Stream_xxx
	SearchLibName *string `json:"SearchLibName,omitempty" xml:"SearchLibName,omitempty"`
}

func (s CreateStreamToSearchLibRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStreamToSearchLibRequest) GoString() string {
	return s.String()
}

func (s *CreateStreamToSearchLibRequest) GetInput() *string {
	return s.Input
}

func (s *CreateStreamToSearchLibRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateStreamToSearchLibRequest) GetSearchLibName() *string {
	return s.SearchLibName
}

func (s *CreateStreamToSearchLibRequest) SetInput(v string) *CreateStreamToSearchLibRequest {
	s.Input = &v
	return s
}

func (s *CreateStreamToSearchLibRequest) SetNamespace(v string) *CreateStreamToSearchLibRequest {
	s.Namespace = &v
	return s
}

func (s *CreateStreamToSearchLibRequest) SetSearchLibName(v string) *CreateStreamToSearchLibRequest {
	s.SearchLibName = &v
	return s
}

func (s *CreateStreamToSearchLibRequest) Validate() error {
	return dara.Validate(s)
}

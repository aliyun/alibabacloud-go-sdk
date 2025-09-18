// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseStreamToSearchLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *CloseStreamToSearchLibRequest
	GetMediaId() *string
	SetNamespace(v string) *CloseStreamToSearchLibRequest
	GetNamespace() *string
	SetSearchLibName(v string) *CloseStreamToSearchLibRequest
	GetSearchLibName() *string
}

type CloseStreamToSearchLibRequest struct {
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// name-1
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// Stream_xxx
	SearchLibName *string `json:"SearchLibName,omitempty" xml:"SearchLibName,omitempty"`
}

func (s CloseStreamToSearchLibRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseStreamToSearchLibRequest) GoString() string {
	return s.String()
}

func (s *CloseStreamToSearchLibRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *CloseStreamToSearchLibRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CloseStreamToSearchLibRequest) GetSearchLibName() *string {
	return s.SearchLibName
}

func (s *CloseStreamToSearchLibRequest) SetMediaId(v string) *CloseStreamToSearchLibRequest {
	s.MediaId = &v
	return s
}

func (s *CloseStreamToSearchLibRequest) SetNamespace(v string) *CloseStreamToSearchLibRequest {
	s.Namespace = &v
	return s
}

func (s *CloseStreamToSearchLibRequest) SetSearchLibName(v string) *CloseStreamToSearchLibRequest {
	s.SearchLibName = &v
	return s
}

func (s *CloseStreamToSearchLibRequest) Validate() error {
	return dara.Validate(s)
}

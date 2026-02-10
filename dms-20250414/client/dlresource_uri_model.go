// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDLResourceUri interface {
	dara.Model
	String() string
	GoString() string
	SetResourceType(v string) *DLResourceUri
	GetResourceType() *string
	SetUri(v string) *DLResourceUri
	GetUri() *string
}

type DLResourceUri struct {
	// example:
	//
	// JAR
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// hdfs://name****.example.com:8020/user/hive/udfs/my_***.jar
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s DLResourceUri) String() string {
	return dara.Prettify(s)
}

func (s DLResourceUri) GoString() string {
	return s.String()
}

func (s *DLResourceUri) GetResourceType() *string {
	return s.ResourceType
}

func (s *DLResourceUri) GetUri() *string {
	return s.Uri
}

func (s *DLResourceUri) SetResourceType(v string) *DLResourceUri {
	s.ResourceType = &v
	return s
}

func (s *DLResourceUri) SetUri(v string) *DLResourceUri {
	s.Uri = &v
	return s
}

func (s *DLResourceUri) Validate() error {
	return dara.Validate(s)
}

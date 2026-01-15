// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDumpMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *DumpMetaRequest
	GetInstanceName() *string
}

type DumpMetaRequest struct {
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// imagesearchName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s DumpMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s DumpMetaRequest) GoString() string {
	return s.String()
}

func (s *DumpMetaRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DumpMetaRequest) SetInstanceName(v string) *DumpMetaRequest {
	s.InstanceName = &v
	return s
}

func (s *DumpMetaRequest) Validate() error {
	return dara.Validate(s)
}

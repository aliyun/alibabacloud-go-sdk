// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOtsDetail interface {
	dara.Model
	String() string
	GoString() string
	SetTableNames(v []*string) *OtsDetail
	GetTableNames() []*string
}

type OtsDetail struct {
	TableNames []*string `json:"TableNames,omitempty" xml:"TableNames,omitempty" type:"Repeated"`
}

func (s OtsDetail) String() string {
	return dara.Prettify(s)
}

func (s OtsDetail) GoString() string {
	return s.String()
}

func (s *OtsDetail) GetTableNames() []*string {
	return s.TableNames
}

func (s *OtsDetail) SetTableNames(v []*string) *OtsDetail {
	s.TableNames = v
	return s
}

func (s *OtsDetail) Validate() error {
	return dara.Validate(s)
}

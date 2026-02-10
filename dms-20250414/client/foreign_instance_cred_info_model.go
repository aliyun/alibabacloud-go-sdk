// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForeignInstanceCredInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCredInfo(v map[string]*string) *ForeignInstanceCredInfo
	GetCredInfo() map[string]*string
	SetCredType(v string) *ForeignInstanceCredInfo
	GetCredType() *string
}

type ForeignInstanceCredInfo struct {
	// The information about the credential.
	CredInfo map[string]*string `json:"CredInfo,omitempty" xml:"CredInfo,omitempty"`
	// The type of the credential. Set the value to DEFAULT.
	//
	// example:
	//
	// DEFAULT
	CredType *string `json:"CredType,omitempty" xml:"CredType,omitempty"`
}

func (s ForeignInstanceCredInfo) String() string {
	return dara.Prettify(s)
}

func (s ForeignInstanceCredInfo) GoString() string {
	return s.String()
}

func (s *ForeignInstanceCredInfo) GetCredInfo() map[string]*string {
	return s.CredInfo
}

func (s *ForeignInstanceCredInfo) GetCredType() *string {
	return s.CredType
}

func (s *ForeignInstanceCredInfo) SetCredInfo(v map[string]*string) *ForeignInstanceCredInfo {
	s.CredInfo = v
	return s
}

func (s *ForeignInstanceCredInfo) SetCredType(v string) *ForeignInstanceCredInfo {
	s.CredType = &v
	return s
}

func (s *ForeignInstanceCredInfo) Validate() error {
	return dara.Validate(s)
}

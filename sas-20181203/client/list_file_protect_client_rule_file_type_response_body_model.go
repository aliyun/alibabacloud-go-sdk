// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileProtectClientRuleFileTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileTypes(v []*string) *ListFileProtectClientRuleFileTypeResponseBody
	GetFileTypes() []*string
	SetRequestId(v string) *ListFileProtectClientRuleFileTypeResponseBody
	GetRequestId() *string
}

type ListFileProtectClientRuleFileTypeResponseBody struct {
	FileTypes []*string `json:"FileTypes,omitempty" xml:"FileTypes,omitempty" type:"Repeated"`
	// example:
	//
	// F8B6F758-BCD4-597A-8A2C-DA5A552C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFileProtectClientRuleFileTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectClientRuleFileTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ListFileProtectClientRuleFileTypeResponseBody) GetFileTypes() []*string {
	return s.FileTypes
}

func (s *ListFileProtectClientRuleFileTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFileProtectClientRuleFileTypeResponseBody) SetFileTypes(v []*string) *ListFileProtectClientRuleFileTypeResponseBody {
	s.FileTypes = v
	return s
}

func (s *ListFileProtectClientRuleFileTypeResponseBody) SetRequestId(v string) *ListFileProtectClientRuleFileTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFileProtectClientRuleFileTypeResponseBody) Validate() error {
	return dara.Validate(s)
}

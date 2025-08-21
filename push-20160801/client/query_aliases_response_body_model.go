// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAliasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliasInfos(v *QueryAliasesResponseBodyAliasInfos) *QueryAliasesResponseBody
	GetAliasInfos() *QueryAliasesResponseBodyAliasInfos
	SetRequestId(v string) *QueryAliasesResponseBody
	GetRequestId() *string
}

type QueryAliasesResponseBody struct {
	AliasInfos *QueryAliasesResponseBodyAliasInfos `json:"AliasInfos,omitempty" xml:"AliasInfos,omitempty" type:"Struct"`
	// example:
	//
	// 159E4422-6624-4750-8943-DFD98D34858C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryAliasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAliasesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAliasesResponseBody) GetAliasInfos() *QueryAliasesResponseBodyAliasInfos {
	return s.AliasInfos
}

func (s *QueryAliasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAliasesResponseBody) SetAliasInfos(v *QueryAliasesResponseBodyAliasInfos) *QueryAliasesResponseBody {
	s.AliasInfos = v
	return s
}

func (s *QueryAliasesResponseBody) SetRequestId(v string) *QueryAliasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAliasesResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryAliasesResponseBodyAliasInfos struct {
	AliasInfo []*QueryAliasesResponseBodyAliasInfosAliasInfo `json:"AliasInfo,omitempty" xml:"AliasInfo,omitempty" type:"Repeated"`
}

func (s QueryAliasesResponseBodyAliasInfos) String() string {
	return dara.Prettify(s)
}

func (s QueryAliasesResponseBodyAliasInfos) GoString() string {
	return s.String()
}

func (s *QueryAliasesResponseBodyAliasInfos) GetAliasInfo() []*QueryAliasesResponseBodyAliasInfosAliasInfo {
	return s.AliasInfo
}

func (s *QueryAliasesResponseBodyAliasInfos) SetAliasInfo(v []*QueryAliasesResponseBodyAliasInfosAliasInfo) *QueryAliasesResponseBodyAliasInfos {
	s.AliasInfo = v
	return s
}

func (s *QueryAliasesResponseBodyAliasInfos) Validate() error {
	return dara.Validate(s)
}

type QueryAliasesResponseBodyAliasInfosAliasInfo struct {
	// example:
	//
	// test_alias1
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
}

func (s QueryAliasesResponseBodyAliasInfosAliasInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryAliasesResponseBodyAliasInfosAliasInfo) GoString() string {
	return s.String()
}

func (s *QueryAliasesResponseBodyAliasInfosAliasInfo) GetAliasName() *string {
	return s.AliasName
}

func (s *QueryAliasesResponseBodyAliasInfosAliasInfo) SetAliasName(v string) *QueryAliasesResponseBodyAliasInfosAliasInfo {
	s.AliasName = &v
	return s
}

func (s *QueryAliasesResponseBodyAliasInfosAliasInfo) Validate() error {
	return dara.Validate(s)
}

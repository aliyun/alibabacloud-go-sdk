// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDBVersionResponseBody
	GetRequestId() *string
	SetVersionSuggestion(v string) *DescribeDBVersionResponseBody
	GetVersionSuggestion() *string
}

type DescribeDBVersionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 25C11EE5-B7E8-481A-A07C-BD619971A570
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The recommended upgrade version in the format of "major version,minor version" (separated by a comma). The first value is the target version for major engine version upgrade, and the second value is the target version for minor engine version update.
	//
	// example:
	//
	// mm.v7.4.2.7-202608031659,mm.v7.3.2.12-202608071438
	VersionSuggestion *string `json:"VersionSuggestion,omitempty" xml:"VersionSuggestion,omitempty"`
}

func (s DescribeDBVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBVersionResponseBody) GetVersionSuggestion() *string {
	return s.VersionSuggestion
}

func (s *DescribeDBVersionResponseBody) SetRequestId(v string) *DescribeDBVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBVersionResponseBody) SetVersionSuggestion(v string) *DescribeDBVersionResponseBody {
	s.VersionSuggestion = &v
	return s
}

func (s *DescribeDBVersionResponseBody) Validate() error {
	return dara.Validate(s)
}

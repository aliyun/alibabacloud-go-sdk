// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *DescribePlayDetailRequest
	GetLanguage() *string
	SetPlayTs(v string) *DescribePlayDetailRequest
	GetPlayTs() *string
	SetSessionId(v string) *DescribePlayDetailRequest
	GetSessionId() *string
}

type DescribePlayDetailRequest struct {
	Language  *string `json:"Language,omitempty" xml:"Language,omitempty"`
	PlayTs    *string `json:"PlayTs,omitempty" xml:"PlayTs,omitempty"`
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s DescribePlayDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribePlayDetailRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribePlayDetailRequest) GetPlayTs() *string {
	return s.PlayTs
}

func (s *DescribePlayDetailRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribePlayDetailRequest) SetLanguage(v string) *DescribePlayDetailRequest {
	s.Language = &v
	return s
}

func (s *DescribePlayDetailRequest) SetPlayTs(v string) *DescribePlayDetailRequest {
	s.PlayTs = &v
	return s
}

func (s *DescribePlayDetailRequest) SetSessionId(v string) *DescribePlayDetailRequest {
	s.SessionId = &v
	return s
}

func (s *DescribePlayDetailRequest) Validate() error {
	return dara.Validate(s)
}

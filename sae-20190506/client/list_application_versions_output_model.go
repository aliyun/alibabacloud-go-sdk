// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationVersionsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *ListApplicationVersionsOutput
	GetDirection() *string
	SetNextToken(v string) *ListApplicationVersionsOutput
	GetNextToken() *string
	SetRequestId(v string) *ListApplicationVersionsOutput
	GetRequestId() *string
	SetVersions(v []*Version) *ListApplicationVersionsOutput
	GetVersions() []*Version
}

type ListApplicationVersionsOutput struct {
	Direction *string    `json:"direction,omitempty" xml:"direction,omitempty"`
	NextToken *string    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId *string    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Versions  []*Version `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s ListApplicationVersionsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationVersionsOutput) GoString() string {
	return s.String()
}

func (s *ListApplicationVersionsOutput) GetDirection() *string {
	return s.Direction
}

func (s *ListApplicationVersionsOutput) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApplicationVersionsOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationVersionsOutput) GetVersions() []*Version {
	return s.Versions
}

func (s *ListApplicationVersionsOutput) SetDirection(v string) *ListApplicationVersionsOutput {
	s.Direction = &v
	return s
}

func (s *ListApplicationVersionsOutput) SetNextToken(v string) *ListApplicationVersionsOutput {
	s.NextToken = &v
	return s
}

func (s *ListApplicationVersionsOutput) SetRequestId(v string) *ListApplicationVersionsOutput {
	s.RequestId = &v
	return s
}

func (s *ListApplicationVersionsOutput) SetVersions(v []*Version) *ListApplicationVersionsOutput {
	s.Versions = v
	return s
}

func (s *ListApplicationVersionsOutput) Validate() error {
	if s.Versions != nil {
		for _, item := range s.Versions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

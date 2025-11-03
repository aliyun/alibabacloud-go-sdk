// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableEngineVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEngineVersions(v *DescribeAvailableEngineVersionResponseBodyEngineVersions) *DescribeAvailableEngineVersionResponseBody
	GetEngineVersions() *DescribeAvailableEngineVersionResponseBodyEngineVersions
	SetRequestId(v string) *DescribeAvailableEngineVersionResponseBody
	GetRequestId() *string
}

type DescribeAvailableEngineVersionResponseBody struct {
	// The list of one or more engine versions to which an ApsaraDB for MongoDB instance can be upgraded.
	//
	// >  An empty string is returned if the latest version is being used.
	EngineVersions *DescribeAvailableEngineVersionResponseBodyEngineVersions `json:"EngineVersions,omitempty" xml:"EngineVersions,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 52507B6B-003B-47A3-A0A3-9FE992C7A243
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAvailableEngineVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableEngineVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableEngineVersionResponseBody) GetEngineVersions() *DescribeAvailableEngineVersionResponseBodyEngineVersions {
	return s.EngineVersions
}

func (s *DescribeAvailableEngineVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAvailableEngineVersionResponseBody) SetEngineVersions(v *DescribeAvailableEngineVersionResponseBodyEngineVersions) *DescribeAvailableEngineVersionResponseBody {
	s.EngineVersions = v
	return s
}

func (s *DescribeAvailableEngineVersionResponseBody) SetRequestId(v string) *DescribeAvailableEngineVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableEngineVersionResponseBody) Validate() error {
	if s.EngineVersions != nil {
		if err := s.EngineVersions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAvailableEngineVersionResponseBodyEngineVersions struct {
	EngineVersion []*string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty" type:"Repeated"`
}

func (s DescribeAvailableEngineVersionResponseBodyEngineVersions) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableEngineVersionResponseBodyEngineVersions) GoString() string {
	return s.String()
}

func (s *DescribeAvailableEngineVersionResponseBodyEngineVersions) GetEngineVersion() []*string {
	return s.EngineVersion
}

func (s *DescribeAvailableEngineVersionResponseBodyEngineVersions) SetEngineVersion(v []*string) *DescribeAvailableEngineVersionResponseBodyEngineVersions {
	s.EngineVersion = v
	return s
}

func (s *DescribeAvailableEngineVersionResponseBodyEngineVersions) Validate() error {
	return dara.Validate(s)
}

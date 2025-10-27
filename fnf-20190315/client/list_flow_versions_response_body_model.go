// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFlowVersions(v []*ListFlowVersionsResponseBodyFlowVersions) *ListFlowVersionsResponseBody
	GetFlowVersions() []*ListFlowVersionsResponseBodyFlowVersions
	SetNextToken(v string) *ListFlowVersionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListFlowVersionsResponseBody
	GetRequestId() *string
}

type ListFlowVersionsResponseBody struct {
	FlowVersions []*ListFlowVersionsResponseBodyFlowVersions `json:"FlowVersions,omitempty" xml:"FlowVersions,omitempty" type:"Repeated"`
	// list token
	//
	// example:
	//
	// token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 294D68C1-5108-5971-853A-1A9CC87B4816
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFlowVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFlowVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowVersionsResponseBody) GetFlowVersions() []*ListFlowVersionsResponseBodyFlowVersions {
	return s.FlowVersions
}

func (s *ListFlowVersionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFlowVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFlowVersionsResponseBody) SetFlowVersions(v []*ListFlowVersionsResponseBodyFlowVersions) *ListFlowVersionsResponseBody {
	s.FlowVersions = v
	return s
}

func (s *ListFlowVersionsResponseBody) SetNextToken(v string) *ListFlowVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListFlowVersionsResponseBody) SetRequestId(v string) *ListFlowVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowVersionsResponseBody) Validate() error {
	if s.FlowVersions != nil {
		for _, item := range s.FlowVersions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFlowVersionsResponseBodyFlowVersions struct {
	// example:
	//
	// 2025-10-24T14:11:26+08:00
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// version description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListFlowVersionsResponseBodyFlowVersions) String() string {
	return dara.Prettify(s)
}

func (s ListFlowVersionsResponseBodyFlowVersions) GoString() string {
	return s.String()
}

func (s *ListFlowVersionsResponseBodyFlowVersions) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListFlowVersionsResponseBodyFlowVersions) GetDescription() *string {
	return s.Description
}

func (s *ListFlowVersionsResponseBodyFlowVersions) GetVersion() *string {
	return s.Version
}

func (s *ListFlowVersionsResponseBodyFlowVersions) SetCreatedTime(v string) *ListFlowVersionsResponseBodyFlowVersions {
	s.CreatedTime = &v
	return s
}

func (s *ListFlowVersionsResponseBodyFlowVersions) SetDescription(v string) *ListFlowVersionsResponseBodyFlowVersions {
	s.Description = &v
	return s
}

func (s *ListFlowVersionsResponseBodyFlowVersions) SetVersion(v string) *ListFlowVersionsResponseBodyFlowVersions {
	s.Version = &v
	return s
}

func (s *ListFlowVersionsResponseBodyFlowVersions) Validate() error {
	return dara.Validate(s)
}

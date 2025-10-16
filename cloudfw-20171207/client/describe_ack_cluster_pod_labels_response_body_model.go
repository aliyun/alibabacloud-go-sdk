// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAckClusterPodLabelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAckLabels(v []*DescribeAckClusterPodLabelsResponseBodyAckLabels) *DescribeAckClusterPodLabelsResponseBody
	GetAckLabels() []*DescribeAckClusterPodLabelsResponseBodyAckLabels
	SetRequestId(v string) *DescribeAckClusterPodLabelsResponseBody
	GetRequestId() *string
}

type DescribeAckClusterPodLabelsResponseBody struct {
	AckLabels []*DescribeAckClusterPodLabelsResponseBodyAckLabels `json:"AckLabels,omitempty" xml:"AckLabels,omitempty" type:"Repeated"`
	// example:
	//
	// 6169C0A4-B91A-5D48-AE4D-B9432D15****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAckClusterPodLabelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAckClusterPodLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAckClusterPodLabelsResponseBody) GetAckLabels() []*DescribeAckClusterPodLabelsResponseBodyAckLabels {
	return s.AckLabels
}

func (s *DescribeAckClusterPodLabelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAckClusterPodLabelsResponseBody) SetAckLabels(v []*DescribeAckClusterPodLabelsResponseBodyAckLabels) *DescribeAckClusterPodLabelsResponseBody {
	s.AckLabels = v
	return s
}

func (s *DescribeAckClusterPodLabelsResponseBody) SetRequestId(v string) *DescribeAckClusterPodLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAckClusterPodLabelsResponseBody) Validate() error {
	if s.AckLabels != nil {
		for _, item := range s.AckLabels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAckClusterPodLabelsResponseBodyAckLabels struct {
	// example:
	//
	// app
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// storage-operator
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAckClusterPodLabelsResponseBodyAckLabels) String() string {
	return dara.Prettify(s)
}

func (s DescribeAckClusterPodLabelsResponseBodyAckLabels) GoString() string {
	return s.String()
}

func (s *DescribeAckClusterPodLabelsResponseBodyAckLabels) GetKey() *string {
	return s.Key
}

func (s *DescribeAckClusterPodLabelsResponseBodyAckLabels) GetValue() *string {
	return s.Value
}

func (s *DescribeAckClusterPodLabelsResponseBodyAckLabels) SetKey(v string) *DescribeAckClusterPodLabelsResponseBodyAckLabels {
	s.Key = &v
	return s
}

func (s *DescribeAckClusterPodLabelsResponseBodyAckLabels) SetValue(v string) *DescribeAckClusterPodLabelsResponseBodyAckLabels {
	s.Value = &v
	return s
}

func (s *DescribeAckClusterPodLabelsResponseBodyAckLabels) Validate() error {
	return dara.Validate(s)
}

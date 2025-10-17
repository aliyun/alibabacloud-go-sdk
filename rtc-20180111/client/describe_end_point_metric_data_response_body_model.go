// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEndPointMetricDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPubMetrics(v []*DescribeEndPointMetricDataResponseBodyPubMetrics) *DescribeEndPointMetricDataResponseBody
	GetPubMetrics() []*DescribeEndPointMetricDataResponseBodyPubMetrics
	SetRequestId(v string) *DescribeEndPointMetricDataResponseBody
	GetRequestId() *string
	SetSubMetrics(v []*DescribeEndPointMetricDataResponseBodySubMetrics) *DescribeEndPointMetricDataResponseBody
	GetSubMetrics() []*DescribeEndPointMetricDataResponseBodySubMetrics
}

type DescribeEndPointMetricDataResponseBody struct {
	PubMetrics []*DescribeEndPointMetricDataResponseBodyPubMetrics `json:"PubMetrics,omitempty" xml:"PubMetrics,omitempty" type:"Repeated"`
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId  *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubMetrics []*DescribeEndPointMetricDataResponseBodySubMetrics `json:"SubMetrics,omitempty" xml:"SubMetrics,omitempty" type:"Repeated"`
}

func (s DescribeEndPointMetricDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndPointMetricDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEndPointMetricDataResponseBody) GetPubMetrics() []*DescribeEndPointMetricDataResponseBodyPubMetrics {
	return s.PubMetrics
}

func (s *DescribeEndPointMetricDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEndPointMetricDataResponseBody) GetSubMetrics() []*DescribeEndPointMetricDataResponseBodySubMetrics {
	return s.SubMetrics
}

func (s *DescribeEndPointMetricDataResponseBody) SetPubMetrics(v []*DescribeEndPointMetricDataResponseBodyPubMetrics) *DescribeEndPointMetricDataResponseBody {
	s.PubMetrics = v
	return s
}

func (s *DescribeEndPointMetricDataResponseBody) SetRequestId(v string) *DescribeEndPointMetricDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEndPointMetricDataResponseBody) SetSubMetrics(v []*DescribeEndPointMetricDataResponseBodySubMetrics) *DescribeEndPointMetricDataResponseBody {
	s.SubMetrics = v
	return s
}

func (s *DescribeEndPointMetricDataResponseBody) Validate() error {
	if s.PubMetrics != nil {
		for _, item := range s.PubMetrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SubMetrics != nil {
		for _, item := range s.SubMetrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEndPointMetricDataResponseBodyPubMetrics struct {
	Nodes []*DescribeEndPointMetricDataResponseBodyPubMetricsNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// VIDEO_STUCK_CAMERA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// testuserid
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeEndPointMetricDataResponseBodyPubMetrics) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndPointMetricDataResponseBodyPubMetrics) GoString() string {
	return s.String()
}

func (s *DescribeEndPointMetricDataResponseBodyPubMetrics) GetNodes() []*DescribeEndPointMetricDataResponseBodyPubMetricsNodes {
	return s.Nodes
}

func (s *DescribeEndPointMetricDataResponseBodyPubMetrics) GetType() *string {
	return s.Type
}

func (s *DescribeEndPointMetricDataResponseBodyPubMetrics) GetUserId() *string {
	return s.UserId
}

func (s *DescribeEndPointMetricDataResponseBodyPubMetrics) SetNodes(v []*DescribeEndPointMetricDataResponseBodyPubMetricsNodes) *DescribeEndPointMetricDataResponseBodyPubMetrics {
	s.Nodes = v
	return s
}

func (s *DescribeEndPointMetricDataResponseBodyPubMetrics) SetType(v string) *DescribeEndPointMetricDataResponseBodyPubMetrics {
	s.Type = &v
	return s
}

func (s *DescribeEndPointMetricDataResponseBodyPubMetrics) SetUserId(v string) *DescribeEndPointMetricDataResponseBodyPubMetrics {
	s.UserId = &v
	return s
}

func (s *DescribeEndPointMetricDataResponseBodyPubMetrics) Validate() error {
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEndPointMetricDataResponseBodyPubMetricsNodes struct {
	Ext map[string]interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// example:
	//
	// 1548670257
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 230100
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeEndPointMetricDataResponseBodyPubMetricsNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndPointMetricDataResponseBodyPubMetricsNodes) GoString() string {
	return s.String()
}

func (s *DescribeEndPointMetricDataResponseBodyPubMetricsNodes) GetExt() map[string]interface{} {
	return s.Ext
}

func (s *DescribeEndPointMetricDataResponseBodyPubMetricsNodes) GetX() *string {
	return s.X
}

func (s *DescribeEndPointMetricDataResponseBodyPubMetricsNodes) GetY() *string {
	return s.Y
}

func (s *DescribeEndPointMetricDataResponseBodyPubMetricsNodes) SetExt(v map[string]interface{}) *DescribeEndPointMetricDataResponseBodyPubMetricsNodes {
	s.Ext = v
	return s
}

func (s *DescribeEndPointMetricDataResponseBodyPubMetricsNodes) SetX(v string) *DescribeEndPointMetricDataResponseBodyPubMetricsNodes {
	s.X = &v
	return s
}

func (s *DescribeEndPointMetricDataResponseBodyPubMetricsNodes) SetY(v string) *DescribeEndPointMetricDataResponseBodyPubMetricsNodes {
	s.Y = &v
	return s
}

func (s *DescribeEndPointMetricDataResponseBodyPubMetricsNodes) Validate() error {
	return dara.Validate(s)
}

type DescribeEndPointMetricDataResponseBodySubMetrics struct {
	Nodes []*DescribeEndPointMetricDataResponseBodySubMetricsNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// VIDEO_STUCK_CAMERA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// testuserid
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeEndPointMetricDataResponseBodySubMetrics) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndPointMetricDataResponseBodySubMetrics) GoString() string {
	return s.String()
}

func (s *DescribeEndPointMetricDataResponseBodySubMetrics) GetNodes() []*DescribeEndPointMetricDataResponseBodySubMetricsNodes {
	return s.Nodes
}

func (s *DescribeEndPointMetricDataResponseBodySubMetrics) GetType() *string {
	return s.Type
}

func (s *DescribeEndPointMetricDataResponseBodySubMetrics) GetUserId() *string {
	return s.UserId
}

func (s *DescribeEndPointMetricDataResponseBodySubMetrics) SetNodes(v []*DescribeEndPointMetricDataResponseBodySubMetricsNodes) *DescribeEndPointMetricDataResponseBodySubMetrics {
	s.Nodes = v
	return s
}

func (s *DescribeEndPointMetricDataResponseBodySubMetrics) SetType(v string) *DescribeEndPointMetricDataResponseBodySubMetrics {
	s.Type = &v
	return s
}

func (s *DescribeEndPointMetricDataResponseBodySubMetrics) SetUserId(v string) *DescribeEndPointMetricDataResponseBodySubMetrics {
	s.UserId = &v
	return s
}

func (s *DescribeEndPointMetricDataResponseBodySubMetrics) Validate() error {
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEndPointMetricDataResponseBodySubMetricsNodes struct {
	Ext map[string]interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// example:
	//
	// 1548670257
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 230100
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeEndPointMetricDataResponseBodySubMetricsNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndPointMetricDataResponseBodySubMetricsNodes) GoString() string {
	return s.String()
}

func (s *DescribeEndPointMetricDataResponseBodySubMetricsNodes) GetExt() map[string]interface{} {
	return s.Ext
}

func (s *DescribeEndPointMetricDataResponseBodySubMetricsNodes) GetX() *string {
	return s.X
}

func (s *DescribeEndPointMetricDataResponseBodySubMetricsNodes) GetY() *string {
	return s.Y
}

func (s *DescribeEndPointMetricDataResponseBodySubMetricsNodes) SetExt(v map[string]interface{}) *DescribeEndPointMetricDataResponseBodySubMetricsNodes {
	s.Ext = v
	return s
}

func (s *DescribeEndPointMetricDataResponseBodySubMetricsNodes) SetX(v string) *DescribeEndPointMetricDataResponseBodySubMetricsNodes {
	s.X = &v
	return s
}

func (s *DescribeEndPointMetricDataResponseBodySubMetricsNodes) SetY(v string) *DescribeEndPointMetricDataResponseBodySubMetricsNodes {
	s.Y = &v
	return s
}

func (s *DescribeEndPointMetricDataResponseBodySubMetricsNodes) Validate() error {
	return dara.Validate(s)
}

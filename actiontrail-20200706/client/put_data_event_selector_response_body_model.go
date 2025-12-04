// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutDataEventSelectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataEventSelectors(v string) *PutDataEventSelectorResponseBody
	GetDataEventSelectors() *string
	SetRequestId(v string) *PutDataEventSelectorResponseBody
	GetRequestId() *string
	SetTrailArn(v string) *PutDataEventSelectorResponseBody
	GetTrailArn() *string
}

type PutDataEventSelectorResponseBody struct {
	// example:
	//
	// [{"EventName":{"Equals":["GetObject","CopyObject","AppendObject"]},"ReadWriteType":"All","ServiceName":"Oss"}]
	DataEventSelectors *string `json:"DataEventSelectors,omitempty" xml:"DataEventSelectors,omitempty"`
	// example:
	//
	// 243E1250-32DA-493B-9347-3C7EEE07****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// acs:actiontrail:cn-shanghai:159498693826****:trail/trail-name
	TrailArn *string `json:"TrailArn,omitempty" xml:"TrailArn,omitempty"`
}

func (s PutDataEventSelectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutDataEventSelectorResponseBody) GoString() string {
	return s.String()
}

func (s *PutDataEventSelectorResponseBody) GetDataEventSelectors() *string {
	return s.DataEventSelectors
}

func (s *PutDataEventSelectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutDataEventSelectorResponseBody) GetTrailArn() *string {
	return s.TrailArn
}

func (s *PutDataEventSelectorResponseBody) SetDataEventSelectors(v string) *PutDataEventSelectorResponseBody {
	s.DataEventSelectors = &v
	return s
}

func (s *PutDataEventSelectorResponseBody) SetRequestId(v string) *PutDataEventSelectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutDataEventSelectorResponseBody) SetTrailArn(v string) *PutDataEventSelectorResponseBody {
	s.TrailArn = &v
	return s
}

func (s *PutDataEventSelectorResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScreenScoreThreadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeScreenScoreThreadResponseBodyData) *DescribeScreenScoreThreadResponseBody
	GetData() *DescribeScreenScoreThreadResponseBodyData
	SetRequestId(v string) *DescribeScreenScoreThreadResponseBody
	GetRequestId() *string
}

type DescribeScreenScoreThreadResponseBody struct {
	// The returned data.
	Data *DescribeScreenScoreThreadResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D03DD0FD-6041-5107-AC00-383E28F1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeScreenScoreThreadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScreenScoreThreadResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScreenScoreThreadResponseBody) GetData() *DescribeScreenScoreThreadResponseBodyData {
	return s.Data
}

func (s *DescribeScreenScoreThreadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScreenScoreThreadResponseBody) SetData(v *DescribeScreenScoreThreadResponseBodyData) *DescribeScreenScoreThreadResponseBody {
	s.Data = v
	return s
}

func (s *DescribeScreenScoreThreadResponseBody) SetRequestId(v string) *DescribeScreenScoreThreadResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScreenScoreThreadResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeScreenScoreThreadResponseBodyData struct {
	// The trends of the scores on the security dashboard.
	SocreThread []*string `json:"SocreThread,omitempty" xml:"SocreThread,omitempty" type:"Repeated"`
	// The dates of the scores on the security dashboard.
	SocreThreadDate []*string `json:"SocreThreadDate,omitempty" xml:"SocreThreadDate,omitempty" type:"Repeated"`
}

func (s DescribeScreenScoreThreadResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeScreenScoreThreadResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeScreenScoreThreadResponseBodyData) GetSocreThread() []*string {
	return s.SocreThread
}

func (s *DescribeScreenScoreThreadResponseBodyData) GetSocreThreadDate() []*string {
	return s.SocreThreadDate
}

func (s *DescribeScreenScoreThreadResponseBodyData) SetSocreThread(v []*string) *DescribeScreenScoreThreadResponseBodyData {
	s.SocreThread = v
	return s
}

func (s *DescribeScreenScoreThreadResponseBodyData) SetSocreThreadDate(v []*string) *DescribeScreenScoreThreadResponseBodyData {
	s.SocreThreadDate = v
	return s
}

func (s *DescribeScreenScoreThreadResponseBodyData) Validate() error {
	return dara.Validate(s)
}

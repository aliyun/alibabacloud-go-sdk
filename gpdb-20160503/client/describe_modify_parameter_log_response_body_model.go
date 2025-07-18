// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModifyParameterLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangelogs(v []*DescribeModifyParameterLogResponseBodyChangelogs) *DescribeModifyParameterLogResponseBody
	GetChangelogs() []*DescribeModifyParameterLogResponseBodyChangelogs
	SetRequestId(v string) *DescribeModifyParameterLogResponseBody
	GetRequestId() *string
}

type DescribeModifyParameterLogResponseBody struct {
	// The queried parameter modification logs.
	Changelogs []*DescribeModifyParameterLogResponseBodyChangelogs `json:"Changelogs,omitempty" xml:"Changelogs,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 7565770E-7C45-462D-BA4A-8A5396F2CAD1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeModifyParameterLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeModifyParameterLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModifyParameterLogResponseBody) GetChangelogs() []*DescribeModifyParameterLogResponseBodyChangelogs {
	return s.Changelogs
}

func (s *DescribeModifyParameterLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeModifyParameterLogResponseBody) SetChangelogs(v []*DescribeModifyParameterLogResponseBodyChangelogs) *DescribeModifyParameterLogResponseBody {
	s.Changelogs = v
	return s
}

func (s *DescribeModifyParameterLogResponseBody) SetRequestId(v string) *DescribeModifyParameterLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeModifyParameterLogResponseBodyChangelogs struct {
	// The effective time.
	//
	// example:
	//
	// 2020-05-05T11:22:22Z
	EffectTime *string `json:"EffectTime,omitempty" xml:"EffectTime,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// testkey
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// Indicates whether the modification takes effect.
	//
	// example:
	//
	// true
	ParameterValid *string `json:"ParameterValid,omitempty" xml:"ParameterValid,omitempty"`
	// The original value of the parameter.
	//
	// example:
	//
	// 100
	ParameterValueAfter *string `json:"ParameterValueAfter,omitempty" xml:"ParameterValueAfter,omitempty"`
	// The new value of the parameter.
	//
	// example:
	//
	// 200
	ParameterValueBefore *string `json:"ParameterValueBefore,omitempty" xml:"ParameterValueBefore,omitempty"`
}

func (s DescribeModifyParameterLogResponseBodyChangelogs) String() string {
	return dara.Prettify(s)
}

func (s DescribeModifyParameterLogResponseBodyChangelogs) GoString() string {
	return s.String()
}

func (s *DescribeModifyParameterLogResponseBodyChangelogs) GetEffectTime() *string {
	return s.EffectTime
}

func (s *DescribeModifyParameterLogResponseBodyChangelogs) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeModifyParameterLogResponseBodyChangelogs) GetParameterValid() *string {
	return s.ParameterValid
}

func (s *DescribeModifyParameterLogResponseBodyChangelogs) GetParameterValueAfter() *string {
	return s.ParameterValueAfter
}

func (s *DescribeModifyParameterLogResponseBodyChangelogs) GetParameterValueBefore() *string {
	return s.ParameterValueBefore
}

func (s *DescribeModifyParameterLogResponseBodyChangelogs) SetEffectTime(v string) *DescribeModifyParameterLogResponseBodyChangelogs {
	s.EffectTime = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBodyChangelogs) SetParameterName(v string) *DescribeModifyParameterLogResponseBodyChangelogs {
	s.ParameterName = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBodyChangelogs) SetParameterValid(v string) *DescribeModifyParameterLogResponseBodyChangelogs {
	s.ParameterValid = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBodyChangelogs) SetParameterValueAfter(v string) *DescribeModifyParameterLogResponseBodyChangelogs {
	s.ParameterValueAfter = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBodyChangelogs) SetParameterValueBefore(v string) *DescribeModifyParameterLogResponseBodyChangelogs {
	s.ParameterValueBefore = &v
	return s
}

func (s *DescribeModifyParameterLogResponseBodyChangelogs) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsNetLevelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeEnsNetLevelResponseBody
	GetCode() *int32
	SetEnsNetLevels(v *DescribeEnsNetLevelResponseBodyEnsNetLevels) *DescribeEnsNetLevelResponseBody
	GetEnsNetLevels() *DescribeEnsNetLevelResponseBodyEnsNetLevels
	SetRequestId(v string) *DescribeEnsNetLevelResponseBody
	GetRequestId() *string
}

type DescribeEnsNetLevelResponseBody struct {
	// The returned service code. A value of 0 indicates that the operation was successful.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The network levels.
	EnsNetLevels *DescribeEnsNetLevelResponseBodyEnsNetLevels `json:"EnsNetLevels,omitempty" xml:"EnsNetLevels,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 86A6D421-A0C7-4C01-8648-47377CA6A2CE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEnsNetLevelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsNetLevelResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetLevelResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeEnsNetLevelResponseBody) GetEnsNetLevels() *DescribeEnsNetLevelResponseBodyEnsNetLevels {
	return s.EnsNetLevels
}

func (s *DescribeEnsNetLevelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEnsNetLevelResponseBody) SetCode(v int32) *DescribeEnsNetLevelResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEnsNetLevelResponseBody) SetEnsNetLevels(v *DescribeEnsNetLevelResponseBodyEnsNetLevels) *DescribeEnsNetLevelResponseBody {
	s.EnsNetLevels = v
	return s
}

func (s *DescribeEnsNetLevelResponseBody) SetRequestId(v string) *DescribeEnsNetLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnsNetLevelResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEnsNetLevelResponseBodyEnsNetLevels struct {
	EnsNetLevel []*DescribeEnsNetLevelResponseBodyEnsNetLevelsEnsNetLevel `json:"EnsNetLevel,omitempty" xml:"EnsNetLevel,omitempty" type:"Repeated"`
}

func (s DescribeEnsNetLevelResponseBodyEnsNetLevels) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsNetLevelResponseBodyEnsNetLevels) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetLevelResponseBodyEnsNetLevels) GetEnsNetLevel() []*DescribeEnsNetLevelResponseBodyEnsNetLevelsEnsNetLevel {
	return s.EnsNetLevel
}

func (s *DescribeEnsNetLevelResponseBodyEnsNetLevels) SetEnsNetLevel(v []*DescribeEnsNetLevelResponseBodyEnsNetLevelsEnsNetLevel) *DescribeEnsNetLevelResponseBodyEnsNetLevels {
	s.EnsNetLevel = v
	return s
}

func (s *DescribeEnsNetLevelResponseBodyEnsNetLevels) Validate() error {
	return dara.Validate(s)
}

type DescribeEnsNetLevelResponseBodyEnsNetLevelsEnsNetLevel struct {
	// The network level. Valid values:
	//
	// 	- Big: greater area.
	//
	// 	- Middle: province.
	//
	// 	- Small: city.
	//
	// example:
	//
	// Big
	EnsNetLevelCode *string `json:"EnsNetLevelCode,omitempty" xml:"EnsNetLevelCode,omitempty"`
}

func (s DescribeEnsNetLevelResponseBodyEnsNetLevelsEnsNetLevel) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsNetLevelResponseBodyEnsNetLevelsEnsNetLevel) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetLevelResponseBodyEnsNetLevelsEnsNetLevel) GetEnsNetLevelCode() *string {
	return s.EnsNetLevelCode
}

func (s *DescribeEnsNetLevelResponseBodyEnsNetLevelsEnsNetLevel) SetEnsNetLevelCode(v string) *DescribeEnsNetLevelResponseBodyEnsNetLevelsEnsNetLevel {
	s.EnsNetLevelCode = &v
	return s
}

func (s *DescribeEnsNetLevelResponseBodyEnsNetLevelsEnsNetLevel) Validate() error {
	return dara.Validate(s)
}

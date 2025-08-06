// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMTSStatisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMTSStatisBySpecList(v *GetMTSStatisResponseBodyMTSStatisBySpecList) *GetMTSStatisResponseBody
	GetMTSStatisBySpecList() *GetMTSStatisResponseBodyMTSStatisBySpecList
	SetRequestId(v string) *GetMTSStatisResponseBody
	GetRequestId() *string
}

type GetMTSStatisResponseBody struct {
	MTSStatisBySpecList *GetMTSStatisResponseBodyMTSStatisBySpecList `json:"MTSStatisBySpecList,omitempty" xml:"MTSStatisBySpecList,omitempty" type:"Struct"`
	RequestId           *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMTSStatisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMTSStatisResponseBody) GoString() string {
	return s.String()
}

func (s *GetMTSStatisResponseBody) GetMTSStatisBySpecList() *GetMTSStatisResponseBodyMTSStatisBySpecList {
	return s.MTSStatisBySpecList
}

func (s *GetMTSStatisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMTSStatisResponseBody) SetMTSStatisBySpecList(v *GetMTSStatisResponseBodyMTSStatisBySpecList) *GetMTSStatisResponseBody {
	s.MTSStatisBySpecList = v
	return s
}

func (s *GetMTSStatisResponseBody) SetRequestId(v string) *GetMTSStatisResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMTSStatisResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMTSStatisResponseBodyMTSStatisBySpecList struct {
	MTSStatisBySpec []*GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpec `json:"MTSStatisBySpec,omitempty" xml:"MTSStatisBySpec,omitempty" type:"Repeated"`
}

func (s GetMTSStatisResponseBodyMTSStatisBySpecList) String() string {
	return dara.Prettify(s)
}

func (s GetMTSStatisResponseBodyMTSStatisBySpecList) GoString() string {
	return s.String()
}

func (s *GetMTSStatisResponseBodyMTSStatisBySpecList) GetMTSStatisBySpec() []*GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpec {
	return s.MTSStatisBySpec
}

func (s *GetMTSStatisResponseBodyMTSStatisBySpecList) SetMTSStatisBySpec(v []*GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpec) *GetMTSStatisResponseBodyMTSStatisBySpecList {
	s.MTSStatisBySpec = v
	return s
}

func (s *GetMTSStatisResponseBodyMTSStatisBySpecList) Validate() error {
	return dara.Validate(s)
}

type GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpec struct {
	MTSStatisDOList *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOList `json:"MTSStatisDOList,omitempty" xml:"MTSStatisDOList,omitempty" type:"Struct"`
	Specification   *string                                                                    `json:"Specification,omitempty" xml:"Specification,omitempty"`
}

func (s GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpec) String() string {
	return dara.Prettify(s)
}

func (s GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpec) GoString() string {
	return s.String()
}

func (s *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpec) GetMTSStatisDOList() *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOList {
	return s.MTSStatisDOList
}

func (s *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpec) GetSpecification() *string {
	return s.Specification
}

func (s *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpec) SetMTSStatisDOList(v *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOList) *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpec {
	s.MTSStatisDOList = v
	return s
}

func (s *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpec) SetSpecification(v string) *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpec {
	s.Specification = &v
	return s
}

func (s *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpec) Validate() error {
	return dara.Validate(s)
}

type GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOList struct {
	MTSStatisDO []*GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOListMTSStatisDO `json:"MTSStatisDO,omitempty" xml:"MTSStatisDO,omitempty" type:"Repeated"`
}

func (s GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOList) String() string {
	return dara.Prettify(s)
}

func (s GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOList) GoString() string {
	return s.String()
}

func (s *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOList) GetMTSStatisDO() []*GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOListMTSStatisDO {
	return s.MTSStatisDO
}

func (s *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOList) SetMTSStatisDO(v []*GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOListMTSStatisDO) *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOList {
	s.MTSStatisDO = v
	return s
}

func (s *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOList) Validate() error {
	return dara.Validate(s)
}

type GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOListMTSStatisDO struct {
	Duration      *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	StatTime      *string `json:"StatTime,omitempty" xml:"StatTime,omitempty"`
	StatTimeUTC   *string `json:"StatTimeUTC,omitempty" xml:"StatTimeUTC,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOListMTSStatisDO) String() string {
	return dara.Prettify(s)
}

func (s GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOListMTSStatisDO) GoString() string {
	return s.String()
}

func (s *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOListMTSStatisDO) GetDuration() *int64 {
	return s.Duration
}

func (s *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOListMTSStatisDO) GetSpecification() *string {
	return s.Specification
}

func (s *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOListMTSStatisDO) GetStatTime() *string {
	return s.StatTime
}

func (s *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOListMTSStatisDO) GetStatTimeUTC() *string {
	return s.StatTimeUTC
}

func (s *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOListMTSStatisDO) GetUserId() *string {
	return s.UserId
}

func (s *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOListMTSStatisDO) SetDuration(v int64) *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOListMTSStatisDO {
	s.Duration = &v
	return s
}

func (s *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOListMTSStatisDO) SetSpecification(v string) *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOListMTSStatisDO {
	s.Specification = &v
	return s
}

func (s *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOListMTSStatisDO) SetStatTime(v string) *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOListMTSStatisDO {
	s.StatTime = &v
	return s
}

func (s *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOListMTSStatisDO) SetStatTimeUTC(v string) *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOListMTSStatisDO {
	s.StatTimeUTC = &v
	return s
}

func (s *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOListMTSStatisDO) SetUserId(v string) *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOListMTSStatisDO {
	s.UserId = &v
	return s
}

func (s *GetMTSStatisResponseBodyMTSStatisBySpecListMTSStatisBySpecMTSStatisDOListMTSStatisDO) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserSecDropResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDrops(v int32) *DescribeDcdnUserSecDropResponseBody
	GetDrops() *int32
	SetMsg(v string) *DescribeDcdnUserSecDropResponseBody
	GetMsg() *string
	SetRequestId(v string) *DescribeDcdnUserSecDropResponseBody
	GetRequestId() *string
	SetUuidStr(v string) *DescribeDcdnUserSecDropResponseBody
	GetUuidStr() *string
}

type DescribeDcdnUserSecDropResponseBody struct {
	// The number of packets that are blocked.
	//
	// example:
	//
	// 233023208
	Drops *int32 `json:"Drops,omitempty" xml:"Drops,omitempty"`
	// Indicates whether the information is found.
	//
	// 	- Found
	//
	// 	- Not Found
	//
	// example:
	//
	// Found
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4A1295C0-7A5C-4F27-8D70-C3A648E7037F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The character string that is concatenated based on the request parameters and is used to locate causes when data is not found.
	//
	// example:
	//
	// 1day10811******6429wafDCDN
	UuidStr *string `json:"UuidStr,omitempty" xml:"UuidStr,omitempty"`
}

func (s DescribeDcdnUserSecDropResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserSecDropResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserSecDropResponseBody) GetDrops() *int32 {
	return s.Drops
}

func (s *DescribeDcdnUserSecDropResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *DescribeDcdnUserSecDropResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnUserSecDropResponseBody) GetUuidStr() *string {
	return s.UuidStr
}

func (s *DescribeDcdnUserSecDropResponseBody) SetDrops(v int32) *DescribeDcdnUserSecDropResponseBody {
	s.Drops = &v
	return s
}

func (s *DescribeDcdnUserSecDropResponseBody) SetMsg(v string) *DescribeDcdnUserSecDropResponseBody {
	s.Msg = &v
	return s
}

func (s *DescribeDcdnUserSecDropResponseBody) SetRequestId(v string) *DescribeDcdnUserSecDropResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnUserSecDropResponseBody) SetUuidStr(v string) *DescribeDcdnUserSecDropResponseBody {
	s.UuidStr = &v
	return s
}

func (s *DescribeDcdnUserSecDropResponseBody) Validate() error {
	return dara.Validate(s)
}

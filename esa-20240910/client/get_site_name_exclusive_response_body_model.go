// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSiteNameExclusiveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v string) *GetSiteNameExclusiveResponseBody
	GetEnable() *string
	SetRequestId(v string) *GetSiteNameExclusiveResponseBody
	GetRequestId() *string
}

type GetSiteNameExclusiveResponseBody struct {
	// Indicates whether site hold is enabled. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 35C66C7B-671H-4297-9187-2C4477247A78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSiteNameExclusiveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSiteNameExclusiveResponseBody) GoString() string {
	return s.String()
}

func (s *GetSiteNameExclusiveResponseBody) GetEnable() *string {
	return s.Enable
}

func (s *GetSiteNameExclusiveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSiteNameExclusiveResponseBody) SetEnable(v string) *GetSiteNameExclusiveResponseBody {
	s.Enable = &v
	return s
}

func (s *GetSiteNameExclusiveResponseBody) SetRequestId(v string) *GetSiteNameExclusiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSiteNameExclusiveResponseBody) Validate() error {
	return dara.Validate(s)
}

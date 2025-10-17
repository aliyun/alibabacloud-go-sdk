// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFingerPrintTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFingerPrintTemplates(v []*DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) *DescribeFingerPrintTemplatesResponseBody
	GetFingerPrintTemplates() []*DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates
	SetRequestId(v string) *DescribeFingerPrintTemplatesResponseBody
	GetRequestId() *string
}

type DescribeFingerPrintTemplatesResponseBody struct {
	// The fingerprint templates.
	FingerPrintTemplates []*DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates `json:"FingerPrintTemplates,omitempty" xml:"FingerPrintTemplates,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 9C1D3FBE-84E1-5ABB-AD98-2003AC71****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFingerPrintTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFingerPrintTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFingerPrintTemplatesResponseBody) GetFingerPrintTemplates() []*DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates {
	return s.FingerPrintTemplates
}

func (s *DescribeFingerPrintTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFingerPrintTemplatesResponseBody) SetFingerPrintTemplates(v []*DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) *DescribeFingerPrintTemplatesResponseBody {
	s.FingerPrintTemplates = v
	return s
}

func (s *DescribeFingerPrintTemplatesResponseBody) SetRequestId(v string) *DescribeFingerPrintTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFingerPrintTemplatesResponseBody) Validate() error {
	if s.FingerPrintTemplates != nil {
		for _, item := range s.FingerPrintTemplates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates struct {
	// The client ID. The system generates a unique ID for each client.
	//
	// example:
	//
	// 2258a3d5-b8f8-4d79-a221-eaecf211****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The time when the template was created.
	//
	// example:
	//
	// 2022-03-13T13:26:29Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the template.
	//
	// example:
	//
	// Finger 1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The user ID.
	//
	// example:
	//
	// liming
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The index.
	//
	// example:
	//
	// 1
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The logon time.
	//
	// example:
	//
	// 2022-03-13T13:26:29Z
	LoginTime *string `json:"LoginTime,omitempty" xml:"LoginTime,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-074949****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
}

func (s DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) String() string {
	return dara.Prettify(s)
}

func (s DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) GoString() string {
	return s.String()
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) GetClientId() *string {
	return s.ClientId
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) GetDescription() *string {
	return s.Description
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) GetIndex() *int64 {
	return s.Index
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) GetLoginTime() *string {
	return s.LoginTime
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) SetClientId(v string) *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates {
	s.ClientId = &v
	return s
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) SetCreationTime(v string) *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates {
	s.CreationTime = &v
	return s
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) SetDescription(v string) *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates {
	s.Description = &v
	return s
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) SetEndUserId(v string) *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates {
	s.EndUserId = &v
	return s
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) SetIndex(v int64) *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates {
	s.Index = &v
	return s
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) SetLoginTime(v string) *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates {
	s.LoginTime = &v
	return s
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) SetOfficeSiteId(v string) *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) Validate() error {
	return dara.Validate(s)
}

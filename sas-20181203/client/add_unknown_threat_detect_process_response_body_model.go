// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUnknownThreatDetectProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIdList(v []*AddUnknownThreatDetectProcessResponseBodyIdList) *AddUnknownThreatDetectProcessResponseBody
	GetIdList() []*AddUnknownThreatDetectProcessResponseBodyIdList
	SetRequestId(v string) *AddUnknownThreatDetectProcessResponseBody
	GetRequestId() *string
}

type AddUnknownThreatDetectProcessResponseBody struct {
	IdList []*AddUnknownThreatDetectProcessResponseBodyIdList `json:"IdList,omitempty" xml:"IdList,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// ADE57832-9666-511C-9A80-B87DE2E8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUnknownThreatDetectProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddUnknownThreatDetectProcessResponseBody) GoString() string {
	return s.String()
}

func (s *AddUnknownThreatDetectProcessResponseBody) GetIdList() []*AddUnknownThreatDetectProcessResponseBodyIdList {
	return s.IdList
}

func (s *AddUnknownThreatDetectProcessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddUnknownThreatDetectProcessResponseBody) SetIdList(v []*AddUnknownThreatDetectProcessResponseBodyIdList) *AddUnknownThreatDetectProcessResponseBody {
	s.IdList = v
	return s
}

func (s *AddUnknownThreatDetectProcessResponseBody) SetRequestId(v string) *AddUnknownThreatDetectProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUnknownThreatDetectProcessResponseBody) Validate() error {
	if s.IdList != nil {
		for _, item := range s.IdList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddUnknownThreatDetectProcessResponseBodyIdList struct {
	// example:
	//
	// 92666883
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s AddUnknownThreatDetectProcessResponseBodyIdList) String() string {
	return dara.Prettify(s)
}

func (s AddUnknownThreatDetectProcessResponseBodyIdList) GoString() string {
	return s.String()
}

func (s *AddUnknownThreatDetectProcessResponseBodyIdList) GetId() *string {
	return s.Id
}

func (s *AddUnknownThreatDetectProcessResponseBodyIdList) SetId(v string) *AddUnknownThreatDetectProcessResponseBodyIdList {
	s.Id = &v
	return s
}

func (s *AddUnknownThreatDetectProcessResponseBodyIdList) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreMediaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForbiddenList(v *RestoreMediaResponseBodyForbiddenList) *RestoreMediaResponseBody
	GetForbiddenList() *RestoreMediaResponseBodyForbiddenList
	SetIgnoredList(v *RestoreMediaResponseBodyIgnoredList) *RestoreMediaResponseBody
	GetIgnoredList() *RestoreMediaResponseBodyIgnoredList
	SetRequestId(v string) *RestoreMediaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RestoreMediaResponseBody
	GetSuccess() *bool
}

type RestoreMediaResponseBody struct {
	// The IDs of the media asset that failed to be processed.
	ForbiddenList *RestoreMediaResponseBodyForbiddenList `json:"ForbiddenList,omitempty" xml:"ForbiddenList,omitempty" type:"Struct"`
	// The IDs of the media assets that failed to be obtained.
	IgnoredList *RestoreMediaResponseBodyIgnoredList `json:"IgnoredList,omitempty" xml:"IgnoredList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 8E70E3F8-E2EE-47BC-4677-379D6F28****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RestoreMediaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestoreMediaResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreMediaResponseBody) GetForbiddenList() *RestoreMediaResponseBodyForbiddenList {
	return s.ForbiddenList
}

func (s *RestoreMediaResponseBody) GetIgnoredList() *RestoreMediaResponseBodyIgnoredList {
	return s.IgnoredList
}

func (s *RestoreMediaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestoreMediaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RestoreMediaResponseBody) SetForbiddenList(v *RestoreMediaResponseBodyForbiddenList) *RestoreMediaResponseBody {
	s.ForbiddenList = v
	return s
}

func (s *RestoreMediaResponseBody) SetIgnoredList(v *RestoreMediaResponseBodyIgnoredList) *RestoreMediaResponseBody {
	s.IgnoredList = v
	return s
}

func (s *RestoreMediaResponseBody) SetRequestId(v string) *RestoreMediaResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestoreMediaResponseBody) SetSuccess(v bool) *RestoreMediaResponseBody {
	s.Success = &v
	return s
}

func (s *RestoreMediaResponseBody) Validate() error {
	return dara.Validate(s)
}

type RestoreMediaResponseBodyForbiddenList struct {
	MediaForbiddenReasonDTO []*RestoreMediaResponseBodyForbiddenListMediaForbiddenReasonDTO `json:"MediaForbiddenReasonDTO,omitempty" xml:"MediaForbiddenReasonDTO,omitempty" type:"Repeated"`
}

func (s RestoreMediaResponseBodyForbiddenList) String() string {
	return dara.Prettify(s)
}

func (s RestoreMediaResponseBodyForbiddenList) GoString() string {
	return s.String()
}

func (s *RestoreMediaResponseBodyForbiddenList) GetMediaForbiddenReasonDTO() []*RestoreMediaResponseBodyForbiddenListMediaForbiddenReasonDTO {
	return s.MediaForbiddenReasonDTO
}

func (s *RestoreMediaResponseBodyForbiddenList) SetMediaForbiddenReasonDTO(v []*RestoreMediaResponseBodyForbiddenListMediaForbiddenReasonDTO) *RestoreMediaResponseBodyForbiddenList {
	s.MediaForbiddenReasonDTO = v
	return s
}

func (s *RestoreMediaResponseBodyForbiddenList) Validate() error {
	return dara.Validate(s)
}

type RestoreMediaResponseBodyForbiddenListMediaForbiddenReasonDTO struct {
	// The ID of the media asset.
	//
	// example:
	//
	// fa10ee70898671edb99f6eb3690d****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The reason for the failure.
	//
	// example:
	//
	// Forbidden.RestoreMedia
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s RestoreMediaResponseBodyForbiddenListMediaForbiddenReasonDTO) String() string {
	return dara.Prettify(s)
}

func (s RestoreMediaResponseBodyForbiddenListMediaForbiddenReasonDTO) GoString() string {
	return s.String()
}

func (s *RestoreMediaResponseBodyForbiddenListMediaForbiddenReasonDTO) GetMediaId() *string {
	return s.MediaId
}

func (s *RestoreMediaResponseBodyForbiddenListMediaForbiddenReasonDTO) GetReason() *string {
	return s.Reason
}

func (s *RestoreMediaResponseBodyForbiddenListMediaForbiddenReasonDTO) SetMediaId(v string) *RestoreMediaResponseBodyForbiddenListMediaForbiddenReasonDTO {
	s.MediaId = &v
	return s
}

func (s *RestoreMediaResponseBodyForbiddenListMediaForbiddenReasonDTO) SetReason(v string) *RestoreMediaResponseBodyForbiddenListMediaForbiddenReasonDTO {
	s.Reason = &v
	return s
}

func (s *RestoreMediaResponseBodyForbiddenListMediaForbiddenReasonDTO) Validate() error {
	return dara.Validate(s)
}

type RestoreMediaResponseBodyIgnoredList struct {
	MediaId []*string `json:"MediaId,omitempty" xml:"MediaId,omitempty" type:"Repeated"`
}

func (s RestoreMediaResponseBodyIgnoredList) String() string {
	return dara.Prettify(s)
}

func (s RestoreMediaResponseBodyIgnoredList) GoString() string {
	return s.String()
}

func (s *RestoreMediaResponseBodyIgnoredList) GetMediaId() []*string {
	return s.MediaId
}

func (s *RestoreMediaResponseBodyIgnoredList) SetMediaId(v []*string) *RestoreMediaResponseBodyIgnoredList {
	s.MediaId = v
	return s
}

func (s *RestoreMediaResponseBodyIgnoredList) Validate() error {
	return dara.Validate(s)
}

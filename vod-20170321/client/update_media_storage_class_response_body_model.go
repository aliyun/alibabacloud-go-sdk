// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaStorageClassResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForbiddenList(v *UpdateMediaStorageClassResponseBodyForbiddenList) *UpdateMediaStorageClassResponseBody
	GetForbiddenList() *UpdateMediaStorageClassResponseBodyForbiddenList
	SetIgnoredList(v *UpdateMediaStorageClassResponseBodyIgnoredList) *UpdateMediaStorageClassResponseBody
	GetIgnoredList() *UpdateMediaStorageClassResponseBodyIgnoredList
	SetRequestId(v string) *UpdateMediaStorageClassResponseBody
	GetRequestId() *string
	SetStatus(v string) *UpdateMediaStorageClassResponseBody
	GetStatus() *string
}

type UpdateMediaStorageClassResponseBody struct {
	// The IDs of the media assets that failed to be processed.
	ForbiddenList *UpdateMediaStorageClassResponseBodyForbiddenList `json:"ForbiddenList,omitempty" xml:"ForbiddenList,omitempty" type:"Struct"`
	// The IDs of the media assets that failed to be obtained.
	IgnoredList *UpdateMediaStorageClassResponseBodyIgnoredList `json:"IgnoredList,omitempty" xml:"IgnoredList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 30FA0B7C-3A81-5449-93CD-ACA149E6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of the task. Valid values:
	//
	// 	- **Processing**
	//
	// 	- **Failed**
	//
	// example:
	//
	// Processing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateMediaStorageClassResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaStorageClassResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMediaStorageClassResponseBody) GetForbiddenList() *UpdateMediaStorageClassResponseBodyForbiddenList {
	return s.ForbiddenList
}

func (s *UpdateMediaStorageClassResponseBody) GetIgnoredList() *UpdateMediaStorageClassResponseBodyIgnoredList {
	return s.IgnoredList
}

func (s *UpdateMediaStorageClassResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMediaStorageClassResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateMediaStorageClassResponseBody) SetForbiddenList(v *UpdateMediaStorageClassResponseBodyForbiddenList) *UpdateMediaStorageClassResponseBody {
	s.ForbiddenList = v
	return s
}

func (s *UpdateMediaStorageClassResponseBody) SetIgnoredList(v *UpdateMediaStorageClassResponseBodyIgnoredList) *UpdateMediaStorageClassResponseBody {
	s.IgnoredList = v
	return s
}

func (s *UpdateMediaStorageClassResponseBody) SetRequestId(v string) *UpdateMediaStorageClassResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMediaStorageClassResponseBody) SetStatus(v string) *UpdateMediaStorageClassResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateMediaStorageClassResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateMediaStorageClassResponseBodyForbiddenList struct {
	MediaForbiddenReasonDTO []*UpdateMediaStorageClassResponseBodyForbiddenListMediaForbiddenReasonDTO `json:"MediaForbiddenReasonDTO,omitempty" xml:"MediaForbiddenReasonDTO,omitempty" type:"Repeated"`
}

func (s UpdateMediaStorageClassResponseBodyForbiddenList) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaStorageClassResponseBodyForbiddenList) GoString() string {
	return s.String()
}

func (s *UpdateMediaStorageClassResponseBodyForbiddenList) GetMediaForbiddenReasonDTO() []*UpdateMediaStorageClassResponseBodyForbiddenListMediaForbiddenReasonDTO {
	return s.MediaForbiddenReasonDTO
}

func (s *UpdateMediaStorageClassResponseBodyForbiddenList) SetMediaForbiddenReasonDTO(v []*UpdateMediaStorageClassResponseBodyForbiddenListMediaForbiddenReasonDTO) *UpdateMediaStorageClassResponseBodyForbiddenList {
	s.MediaForbiddenReasonDTO = v
	return s
}

func (s *UpdateMediaStorageClassResponseBodyForbiddenList) Validate() error {
	return dara.Validate(s)
}

type UpdateMediaStorageClassResponseBodyForbiddenListMediaForbiddenReasonDTO struct {
	// The ID of the media asset.
	//
	// example:
	//
	// 19e231ee6e3e417fbf2e92ec2680****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The reason for the failure.
	//
	// example:
	//
	// TargetStorageClassInconsistent
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s UpdateMediaStorageClassResponseBodyForbiddenListMediaForbiddenReasonDTO) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaStorageClassResponseBodyForbiddenListMediaForbiddenReasonDTO) GoString() string {
	return s.String()
}

func (s *UpdateMediaStorageClassResponseBodyForbiddenListMediaForbiddenReasonDTO) GetMediaId() *string {
	return s.MediaId
}

func (s *UpdateMediaStorageClassResponseBodyForbiddenListMediaForbiddenReasonDTO) GetReason() *string {
	return s.Reason
}

func (s *UpdateMediaStorageClassResponseBodyForbiddenListMediaForbiddenReasonDTO) SetMediaId(v string) *UpdateMediaStorageClassResponseBodyForbiddenListMediaForbiddenReasonDTO {
	s.MediaId = &v
	return s
}

func (s *UpdateMediaStorageClassResponseBodyForbiddenListMediaForbiddenReasonDTO) SetReason(v string) *UpdateMediaStorageClassResponseBodyForbiddenListMediaForbiddenReasonDTO {
	s.Reason = &v
	return s
}

func (s *UpdateMediaStorageClassResponseBodyForbiddenListMediaForbiddenReasonDTO) Validate() error {
	return dara.Validate(s)
}

type UpdateMediaStorageClassResponseBodyIgnoredList struct {
	MediaId []*string `json:"MediaId,omitempty" xml:"MediaId,omitempty" type:"Repeated"`
}

func (s UpdateMediaStorageClassResponseBodyIgnoredList) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaStorageClassResponseBodyIgnoredList) GoString() string {
	return s.String()
}

func (s *UpdateMediaStorageClassResponseBodyIgnoredList) GetMediaId() []*string {
	return s.MediaId
}

func (s *UpdateMediaStorageClassResponseBodyIgnoredList) SetMediaId(v []*string) *UpdateMediaStorageClassResponseBodyIgnoredList {
	s.MediaId = v
	return s
}

func (s *UpdateMediaStorageClassResponseBodyIgnoredList) Validate() error {
	return dara.Validate(s)
}

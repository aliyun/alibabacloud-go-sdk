// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLocationDateCluster interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v *Address) *LocationDateCluster
	GetAddress() *Address
	SetClusterId(v string) *LocationDateCluster
	GetClusterId() *string
	SetCreatedAt(v string) *LocationDateCluster
	GetCreatedAt() *string
	SetCustomLabels(v map[string]*string) *LocationDateCluster
	GetCustomLabels() map[string]*string
	SetDriveId(v string) *LocationDateCluster
	GetDriveId() *string
	SetEndTime(v string) *LocationDateCluster
	GetEndTime() *string
	SetLevel(v string) *LocationDateCluster
	GetLevel() *string
	SetStartTime(v string) *LocationDateCluster
	GetStartTime() *string
	SetTitle(v string) *LocationDateCluster
	GetTitle() *string
	SetUpdatedAt(v string) *LocationDateCluster
	GetUpdatedAt() *string
}

type LocationDateCluster struct {
	Address      *Address           `json:"address,omitempty" xml:"address,omitempty"`
	ClusterId    *string            `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	CreatedAt    *string            `json:"created_at,omitempty" xml:"created_at,omitempty"`
	CustomLabels map[string]*string `json:"custom_labels,omitempty" xml:"custom_labels,omitempty"`
	DriveId      *string            `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	EndTime      *string            `json:"end_time,omitempty" xml:"end_time,omitempty"`
	Level        *string            `json:"level,omitempty" xml:"level,omitempty"`
	StartTime    *string            `json:"start_time,omitempty" xml:"start_time,omitempty"`
	Title        *string            `json:"title,omitempty" xml:"title,omitempty"`
	UpdatedAt    *string            `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s LocationDateCluster) String() string {
	return dara.Prettify(s)
}

func (s LocationDateCluster) GoString() string {
	return s.String()
}

func (s *LocationDateCluster) GetAddress() *Address {
	return s.Address
}

func (s *LocationDateCluster) GetClusterId() *string {
	return s.ClusterId
}

func (s *LocationDateCluster) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *LocationDateCluster) GetCustomLabels() map[string]*string {
	return s.CustomLabels
}

func (s *LocationDateCluster) GetDriveId() *string {
	return s.DriveId
}

func (s *LocationDateCluster) GetEndTime() *string {
	return s.EndTime
}

func (s *LocationDateCluster) GetLevel() *string {
	return s.Level
}

func (s *LocationDateCluster) GetStartTime() *string {
	return s.StartTime
}

func (s *LocationDateCluster) GetTitle() *string {
	return s.Title
}

func (s *LocationDateCluster) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *LocationDateCluster) SetAddress(v *Address) *LocationDateCluster {
	s.Address = v
	return s
}

func (s *LocationDateCluster) SetClusterId(v string) *LocationDateCluster {
	s.ClusterId = &v
	return s
}

func (s *LocationDateCluster) SetCreatedAt(v string) *LocationDateCluster {
	s.CreatedAt = &v
	return s
}

func (s *LocationDateCluster) SetCustomLabels(v map[string]*string) *LocationDateCluster {
	s.CustomLabels = v
	return s
}

func (s *LocationDateCluster) SetDriveId(v string) *LocationDateCluster {
	s.DriveId = &v
	return s
}

func (s *LocationDateCluster) SetEndTime(v string) *LocationDateCluster {
	s.EndTime = &v
	return s
}

func (s *LocationDateCluster) SetLevel(v string) *LocationDateCluster {
	s.Level = &v
	return s
}

func (s *LocationDateCluster) SetStartTime(v string) *LocationDateCluster {
	s.StartTime = &v
	return s
}

func (s *LocationDateCluster) SetTitle(v string) *LocationDateCluster {
	s.Title = &v
	return s
}

func (s *LocationDateCluster) SetUpdatedAt(v string) *LocationDateCluster {
	s.UpdatedAt = &v
	return s
}

func (s *LocationDateCluster) Validate() error {
	if s.Address != nil {
		if err := s.Address.Validate(); err != nil {
			return err
		}
	}
	return nil
}

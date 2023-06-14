// Code generated by go-swagger; DO NOT EDIT.

package clustersourcesecret

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/swag"

	objectmetamodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/objectmeta"
)

// VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecret SourceSecret represents a credential used to authenticate to a fluxcd source such as GitRepository.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.fluxcd.sourcesecret.SourceSecret
type VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecret struct {

	// Full name for the Source Secret.
	FullName *VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretFullName `json:"fullName,omitempty"`

	// Metadata for the Source Secret object.
	Meta *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectMeta `json:"meta,omitempty"`

	// Spec for the Source Secret.
	Spec *VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretSpec `json:"spec,omitempty"`

	// Status for the Source Secret.
	Status *VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecretStatus `json:"status,omitempty"`

	// Metadata describing the type of the resource.
	Type *objectmetamodel.VmwareTanzuCoreV1alpha1ObjectType `json:"type,omitempty"`
}

// MarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecret) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecret) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterFluxcdSourcesecret
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

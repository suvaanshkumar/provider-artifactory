package artifactory_local_generic_repository

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("artifactory_local_generic_repository", func(r *config.Resource) {
		
		r.ShortGroup = "artifactory_local_generic_repository"

		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.SetIdentifierArgumentFn = func(base map[string]any, externalName string) {
			base["key"] = externalName
		}

	})
}
